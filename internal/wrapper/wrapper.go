package wrapper

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/Lucino772/envelop/internal/utils"
	"github.com/Lucino772/envelop/pkg/pubsub"
	"github.com/go-cmd/cmd"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Options struct {
	Program  string
	Args     []string
	Dir      string
	Env      []string
	Graceful struct {
		Timeout time.Duration
		Stopper Stopper
	}
	Tasks         []Task
	Services      []Service
	ConfigParsers []ConfigParser
	Logger        *slog.Logger
}

func Run(ctx context.Context, options *Options) error {
	logger := options.Logger

	// TODO: Ensure defaults are set in options
	if options.Graceful.Timeout == 0 {
		options.Graceful.Timeout = 30 * time.Second
	}
	if options.Tasks == nil {
		options.Tasks = make([]Task, 0)
	}
	if options.Services == nil {
		options.Services = make([]Service, 0)
	}
	if options.Dir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		options.Dir = cwd
	}

	// Initialize states
	states, err := NewStates()
	if err != nil {
		return err
	}
	states.Set(&ProcessStatusState{
		Description: "Unknown",
	})
	states.Set(&PlayerState{
		Count:   0,
		Max:     0,
		Players: []string{},
	})

	// Event producer
	eventProducer := pubsub.NewProducer(5, states.HandleEvent)
	options.Tasks = append(
		options.Tasks,
		NewNamedTask(
			"events-producer",
			func(ctx context.Context, _ Wrapper) error {
				return eventProducer.Run(ctx)
			},
		),
	)

	// Prepare command
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return err
	}
	command := cmd.NewCmdOptions(cmd.Options{
		Buffered:  false,
		Streaming: true,
		BeforeExec: []func(cmd *exec.Cmd){
			func(cmd *exec.Cmd) {
				cmd.SysProcAttr.CreationFlags = syscall.CREATE_NEW_PROCESS_GROUP
			},
		},
	}, options.Program, options.Args...)
	command.Dir = options.Dir
	command.Env = append(command.Env, os.Environ()...)
	command.Env = append(command.Env, options.Env...)

	// Prepare wrapper context
	wrapperCtx := WrapperContext{
		dir:            options.Dir,
		stdin:          stdinWriter,
		command:        command,
		eventsProducer: eventProducer,
		logger:         logger,
		states:         states,
		configParsers:  options.ConfigParsers,
	}

	// Run
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		mainErrGroup, mainCtx = errgroup.WithContext(runCtx)
		taskErrGroup, _       = errgroup.WithContext(mainCtx)
	)
	mainErrGroup.SetLimit(2)
	taskErrGroup.SetLimit(-1)

	for _, task := range options.Tasks {
		taskErrGroup.Go(makeRecoverableTask(mainCtx, task, logger, &wrapperCtx))
	}

	mainErrGroup.Go(func() error {
		defer cancel()
		logger.LogAttrs(mainCtx, slog.LevelInfo, "Starting gRPC server")
		err := runGrpcServer(mainCtx, &wrapperCtx, options)
		if err != nil {
			logger.LogAttrs(
				mainCtx,
				slog.LevelError,
				"gRPC server error",
				slog.Any("error", err),
			)
		} else {
			logger.LogAttrs(mainCtx, slog.LevelInfo, "gRPC server stopped")
		}
		return err
	})
	mainErrGroup.Go(func() error {
		defer cancel()
		logger.LogAttrs(mainCtx, slog.LevelInfo, "Starting process")
		err := runProcess(mainCtx, options, &wrapperCtx, stdinReader)
		if err != nil {
			logger.LogAttrs(
				mainCtx,
				slog.LevelError,
				"Process error",
				slog.Any("error", err),
			)
		} else {
			logger.LogAttrs(mainCtx, slog.LevelInfo, "Process stopped")
		}
		return err
	})

	err = mainErrGroup.Wait()
	taskErrGroup.Wait()
	return err
}

func runProcess(ctx context.Context, options *Options, wp *WrapperContext, stdinReader io.Reader) error {
	defer wp.stdin.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	statusChan := wp.command.StartWithStdin(stdinReader)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case value, ok := <-wp.command.Stdout:
				if !ok {
					return
				}
				wp.EmitEvent(ProcessLogEvent{Message: value})
			case value, ok := <-wp.command.Stderr:
				if !ok {
					return
				}
				wp.EmitEvent(ProcessLogEvent{Message: value})
			case <-ctx.Done():
				return
			}
		}
	}()

	var err error = nil
	select {
	case <-ctx.Done():
		gracefulStop(statusChan, options, wp)
		err = wp.command.Stop()
	case <-signalChan:
		gracefulStop(statusChan, options, wp)
		err = wp.command.Stop()
	case status := <-statusChan:
		signal.Stop(signalChan)
		err = status.Error
	}
	wg.Wait()
	return err
}

func gracefulStop(statusChan <-chan cmd.Status, options *Options, wp *WrapperContext) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), options.Graceful.Timeout)
	defer cancel()

	err := options.Graceful.Stopper(wp)
	if err != nil {
		return false, err
	}

	select {
	case <-statusChan:
		return true, nil
	case <-ctx.Done():
		return false, nil
	}
}

type defaultGrpcWrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *defaultGrpcWrappedStream) Context() context.Context {
	return w.ctx
}

func runGrpcServer(ctx context.Context, wp *WrapperContext, options *Options) error {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				return handler(WithWrapper(ctx, wp), req)
			},
		),
		grpc.StreamInterceptor(
			func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				return handler(srv, &defaultGrpcWrappedStream{ss, WithWrapper(ss.Context(), wp)})
			},
		),
	)
	for _, service := range options.Services {
		service.Register(grpcServer)
	}

	quit := make(chan error)
	go func() {
		defer close(quit)
		err := grpcServer.Serve(lis)
		if err != nil {
			quit <- err
		}
	}()

	select {
	case <-ctx.Done():
		grpcServer.Stop()
		return nil
	case err := <-quit:
		return err
	}
}

type WrapperContext struct {
	dir            string
	stdin          io.WriteCloser
	command        *cmd.Cmd
	eventsProducer pubsub.Producer[Event]
	logger         *slog.Logger
	states         *States
	configParsers  []ConfigParser
	config         keyValue
}

func (wp *WrapperContext) Files() fs.FS {
	return os.DirFS(wp.dir)
}

func (wp *WrapperContext) WriteStdin(command string) error {
	_, err := wp.stdin.Write([]byte(fmt.Sprintf("%s\n", command)))
	return err
}

func (wp *WrapperContext) SendSignal(signal os.Signal) error {
	// TODO: Check process started
	pid := wp.command.Status().PID
	if runtime.GOOS == "windows" {
		d, err := syscall.LoadDLL("kernel32.dll")
		if err != nil {
			return err
		}
		p, err := d.FindProc("GenerateConsoleCtrlEvent")
		if err != nil {
			return err
		}
		r, _, err := p.Call(syscall.CTRL_BREAK_EVENT, uintptr(pid))
		if r == 0 {
			return err
		}
		return nil
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Signal(signal)
}

func (wp *WrapperContext) EmitEvent(event any) {
	wp.eventsProducer.Emit(Event{
		Timestamp: time.Now().Unix(),
		Name:      GetEventName(event),
		Data:      event,
	})
}

func (wp *WrapperContext) SubscribeEvents() pubsub.Subscriber[Event] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (Event, bool) {
		return e, true
	})
}

func (wp *WrapperContext) SubscribeLogs() pubsub.Subscriber[string] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (string, bool) {
		if event, ok := e.Data.(ProcessLogEvent); ok {
			return event.Message, true
		}
		return "", false
	})
}

func (wp *WrapperContext) SubscribeStates() pubsub.Subscriber[any] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (any, bool) {
		if event, ok := e.Data.(StateUpdateEvent); ok {
			return event.Data, true
		}
		return nil, false
	})
}

func (wp *WrapperContext) UpdateState(state any) {
	wp.EmitEvent(StateUpdateEvent{
		Name: GetStateName(state),
		Data: state,
	})
}

func (wp *WrapperContext) ReadState(state any) bool {
	return wp.states.Get(state)
}

func (wp *WrapperContext) Logger() *slog.Logger {
	return wp.logger
}

func (wp *WrapperContext) GetServerConfig() (KeyValue, error) {
	if wp.config == nil {
		config := make(keyValue)
		for _, parser := range wp.configParsers {
			if err := parser.Parse(config, wp); err != nil {
				return nil, err
			}
		}
		wp.config = config
	}
	return wp.config, nil
}

type States struct {
	store       map[string]any
	idGenerator func() (string, error)
}

func NewStates() (*States, error) {
	idGenerator, err := utils.NewNanoIDGenerator()
	if err != nil {
		return nil, err
	}
	return &States{
		store:       make(map[string]any),
		idGenerator: idGenerator,
	}, nil
}

func (s *States) HandleEvent(event Event) (Event, bool) {
	id, err := s.idGenerator()
	if err == nil {
		event.Id = id
	}

	if stateEvent, ok := event.Data.(StateUpdateEvent); ok {
		return event, s.Set(stateEvent.Data)
	}
	return event, true
}

func (s *States) Set(state any) bool {
	name := GetStateName(state)
	current, ok := s.store[name]
	var updated bool = false
	if !ok {
		s.store[name] = state
		updated = true
	} else if !reflect.DeepEqual(current, state) {
		s.store[name] = state
		updated = true
	}
	return updated
}

func (s *States) Get(state any) bool {
	if state == nil {
		return false
	}

	value, ok := s.store[GetStateName(state)]
	if !ok {
		return false
	}

	valuePtr := reflect.ValueOf(value)
	if valuePtr.Kind() != reflect.Ptr {
		return false
	}
	reflect.ValueOf(state).Elem().Set(valuePtr.Elem())
	return true
}
