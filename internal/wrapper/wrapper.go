package wrapper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"

	"github.com/Lucino772/envelop/internal/utils"
	"github.com/Lucino772/envelop/pkg/pubsub"
	"github.com/go-cmd/cmd"
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

	// Prepare command
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return err
	}
	command := cmd.NewCmdOptions(cmd.Options{
		Buffered:   false,
		Streaming:  true,
		BeforeExec: []func(cmd *exec.Cmd){setProcessGroupID},
	}, options.Program, options.Args...)
	command.Dir = options.Dir
	command.Env = append(command.Env, os.Environ()...)
	command.Env = append(command.Env, options.Env...)

	// Prepare wrapper context
	wrapperCtx := WrapperContext{
		dir:           options.Dir,
		stdin:         stdinWriter,
		command:       command,
		logger:        logger,
		configParsers: options.ConfigParsers,
		state: ServerState{
			Status: ServerState_Status{
				Description: "Unknown",
			},
			Players: ServerState_Players{
				Count: 0,
				Max:   0,
				List:  make([]ServerState_Player, 0),
			},
		},
		idGenerator: utils.NewNanoIDGenerator(),
	}
	wrapperCtx.eventsProducer = pubsub.NewProducer(100, wrapperCtx.handleEvent)
	options.Tasks = append(
		options.Tasks,
		NewNamedTask(
			"events-producer",
			func(ctx context.Context, _ Wrapper) error {
				return wrapperCtx.eventsProducer.Run(ctx)
			},
		),
		NewNamedTask(
			"grpc-server",
			func(ctx context.Context, wp Wrapper) error {
				return runGrpcServer(ctx, wp, options)
			},
		),
	)

	// Run
	processCtx, cancelProcess := context.WithCancel(ctx)
	defer cancelProcess()
	tasksCtx, cancelTasks := context.WithCancel(context.Background())
	defer cancelTasks()

	pool := NewPool()
	for _, task := range options.Tasks {
		pool.Submit(task, func() error {
			logger.LogAttrs(tasksCtx, slog.LevelInfo, "task started", slog.String("task", task.Name()))
			return task.Run(tasksCtx, &wrapperCtx)
		})
	}
	pool.Submit("process", func() error {
		logger.Log(processCtx, slog.LevelInfo, "process started")
		return runProcess(processCtx, options, &wrapperCtx, stdinReader)
	})
	pool.Monitor()

	var processErr error
	for result := range pool.Results() {
		switch value := result.Key.(type) {
		case Task:
			if result.Error != nil && !errors.Is(result.Error, context.Canceled) {
				logger.LogAttrs(
					tasksCtx,
					slog.LevelError,
					"task stopped with error, retrying",
					slog.String("task", value.Name()),
					slog.Any("error", result.Error),
				)
				pool.Submit(value, func() error {
					logger.LogAttrs(tasksCtx, slog.LevelInfo, "task started", slog.String("task", value.Name()))
					return value.Run(tasksCtx, &wrapperCtx)
				})
			} else {
				logger.LogAttrs(
					tasksCtx,
					slog.LevelInfo,
					"task is done",
					slog.String("task", value.Name()),
				)
			}
		case string:
			if result.Error != nil && !errors.Is(result.Error, context.Canceled) {
				logger.LogAttrs(
					processCtx,
					slog.LevelError,
					"process stopped with error",
					slog.Any("error", err),
				)
				processErr = result.Error
			} else {
				logger.LogAttrs(processCtx, slog.LevelInfo, "process is done")
			}
			cancelTasks()
		}
	}
	<-pool.Done()
	return processErr
}

func runProcess(ctx context.Context, options *Options, wp *WrapperContext, stdinReader io.Reader) error {
	defer wp.stdin.Close()

	signalCtx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

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
	case <-signalCtx.Done():
		gracefulStop(statusChan, options, wp)
		err = wp.command.Stop()
	case status := <-statusChan:
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

func runGrpcServer(ctx context.Context, wp Wrapper, options *Options) error {
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
	configParsers  []ConfigParser
	config         typeConversionMap
	state          ServerState
	idGenerator    func() (string, error)
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
	return sendSignal(wp.command.Status().PID, signal)
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

func (wp *WrapperContext) SubscribeStates() pubsub.Subscriber[ServerState] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (ServerState, bool) {
		if event, ok := e.Data.(StateUpdateEvent); ok {
			return event.State, true
		}
		return ServerState{}, false
	})
}

func (wp *WrapperContext) UpdateState(updateFn func(ServerState) ServerState) {
	state := updateFn(wp.state)
	wp.EmitEvent(StateUpdateEvent{State: state})
}

func (wp *WrapperContext) GetState() ServerState {
	return wp.state
}

func (wp *WrapperContext) Logger() *slog.Logger {
	return wp.logger
}

func (wp *WrapperContext) GetServerConfig() (Map, error) {
	if wp.config == nil {
		config := make(typeConversionMap)
		for _, parser := range wp.configParsers {
			if err := parser.Parse(config, wp); err != nil {
				return nil, err
			}
		}
		wp.config = config
	}
	return wp.config, nil
}

func (wp *WrapperContext) handleEvent(event Event) (Event, bool) {
	id, err := wp.idGenerator()
	if err == nil {
		event.Id = id
	}

	if stateEvent, ok := event.Data.(StateUpdateEvent); ok {
		var updated bool = false
		if !reflect.DeepEqual(wp.state, stateEvent.State) {
			updated = true
			wp.state = stateEvent.State
		}
		return event, updated
	}
	return event, true
}
