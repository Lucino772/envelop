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
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"

	"github.com/Lucino772/envelop/internal/utils"
	"github.com/Lucino772/envelop/pkg/pubsub"
	"github.com/go-cmd/cmd"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type OptFunc func(*wrapper)

type wrapper struct {
	cmd             *cmd.Cmd
	stdinReader     io.Reader
	stdinWriter     io.WriteCloser
	gracefulTimeout time.Duration
	gracefulStopper Stopper
	services        []Service
	tasks           []Task
	eventsProducer  pubsub.Producer[Event]
	states          map[string]any
	idGenerator     func() (string, error)
	logger          *slog.Logger
}

func New(program string, args []string, logger *slog.Logger, opts ...OptFunc) (func(context.Context) error, error) {
	idGenerator, err := utils.NewNanoIDGenerator()
	if err != nil {
		return nil, err
	}
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	command := cmd.NewCmdOptions(cmd.Options{
		Buffered:  false,
		Streaming: true,
	}, program, args...)
	command.Env = append(command.Env, os.Environ()...)

	wp := &wrapper{
		gracefulTimeout: 30 * time.Second,
		services:        make([]Service, 0),
		tasks:           make([]Task, 0),
		cmd:             command,
		stdinReader:     stdinReader,
		stdinWriter:     stdinWriter,
		idGenerator:     idGenerator,
		states:          make(map[string]any),
		logger:          logger,
	}
	wp.eventsProducer = pubsub.NewProducer(5, wp.processEvent)
	wp.setState(&ProcessStatusState{
		Description: "Unknown",
	})
	wp.setState(&PlayerState{
		Count:   0,
		Max:     0,
		Players: []string{},
	})

	for _, opt := range opts {
		opt(wp)
	}
	return wp.Run, nil
}

func (wp *wrapper) Run(parent context.Context) error {
	logger := wp.Logger()

	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	wp.tasks = append(
		wp.tasks,
		NewNamedTask(
			"events-producer",
			func(ctx context.Context, _ Wrapper) error {
				err := wp.eventsProducer.Run(ctx)
				if errors.Is(err, context.Canceled) {
					return nil
				}
				return err
			},
		),
	)

	var (
		mainErrGroup, mainCtx = errgroup.WithContext(ctx)
		taskErrGroup, _       = errgroup.WithContext(mainCtx)
	)
	mainErrGroup.SetLimit(2)
	taskErrGroup.SetLimit(-1)

	for _, task := range wp.tasks {
		task := task
		taskErrGroup.Go(func() error {
			logger := logger.With(slog.String("task", task.Name()))
			logger.LogAttrs(mainCtx, LevelInfo, "Starting task")
			err := task.Run(mainCtx, wp)
			if err != nil {
				logger.LogAttrs(
					mainCtx,
					LevelError,
					"Task error",
					slog.Any("error", err),
				)
			} else {
				logger.LogAttrs(mainCtx, LevelInfo, "Task done")
			}
			return err
		})
	}

	mainErrGroup.Go(func() error {
		defer cancel()
		logger.LogAttrs(mainCtx, LevelInfo, "Starting gRPC server")
		err := wp.runGrpcServer(mainCtx)
		if err != nil {
			logger.LogAttrs(
				mainCtx,
				LevelError,
				"gRPC server error",
				slog.Any("error", err),
			)
		} else {
			logger.LogAttrs(mainCtx, LevelInfo, "gRPC server stopped")
		}
		return err
	})
	mainErrGroup.Go(func() error {
		defer cancel()
		logger.LogAttrs(mainCtx, LevelInfo, "Starting process")
		err := wp.runProcess(mainCtx)
		if err != nil {
			logger.LogAttrs(
				mainCtx,
				LevelError,
				"Process error",
				slog.Any("error", err),
			)
		} else {
			logger.LogAttrs(mainCtx, LevelInfo, "Process stopped")
		}
		return err
	})

	err := mainErrGroup.Wait()
	taskErrGroup.Wait()
	return err
}

func (wp *wrapper) Files() fs.FS {
	return os.DirFS(wp.cmd.Dir)
}

func (wp *wrapper) WriteStdin(command string) error {
	_, err := wp.stdinWriter.Write([]byte(fmt.Sprintf("%s\n", command)))
	return err
}

func (wp *wrapper) SendSignal(signal os.Signal) error {
	process, err := os.FindProcess(wp.cmd.Status().PID)
	if err != nil {
		return err
	}
	return process.Signal(signal)
}

func (wp *wrapper) SubscribeLogs() pubsub.Subscriber[string] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (string, bool) {
		if event, ok := e.Data.(ProcessLogEvent); ok {
			return event.Message, true
		}
		return "", false
	})
}

func (wp *wrapper) SubscribeEvents() pubsub.Subscriber[Event] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (Event, bool) {
		return e, true
	})
}

func (wp *wrapper) EmitEvent(event any) {
	wp.eventsProducer.Emit(Event{
		Timestamp: time.Now().Unix(),
		Name:      GetEventName(event),
		Data:      event,
	})
}

func (wp *wrapper) ReadState(state any) bool {
	if state == nil {
		return false
	}

	value, ok := wp.states[GetStateName(state)]
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

func (wp *wrapper) SubscribeStates() pubsub.Subscriber[any] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (any, bool) {
		if event, ok := e.Data.(StateUpdateEvent); ok {
			return event.Data, true
		}
		return nil, false
	})
}

func (wp *wrapper) UpdateState(state any) {
	wp.EmitEvent(StateUpdateEvent{
		Name: GetStateName(state),
		Data: state,
	})
}

func (wp *wrapper) Logger() *slog.Logger {
	return wp.logger
}

func (w *wrapper) processEvent(event Event) (Event, bool) {
	id, err := w.idGenerator()
	if err == nil {
		event.Id = id
	}

	if stateEvent, ok := event.Data.(StateUpdateEvent); ok {
		return event, w.setState(stateEvent.Data)
	}
	return event, true
}

func (w *wrapper) setState(state any) bool {
	name := GetStateName(state)
	current, ok := w.states[name]

	var updated bool = false
	if !ok {
		w.states[name] = state
		updated = true
	} else if !reflect.DeepEqual(current, state) {
		w.states[name] = state
		updated = true
	}
	return updated
}

func (wp *wrapper) runProcess(ctx context.Context) error {
	defer wp.stdinWriter.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	statusChan := wp.cmd.StartWithStdin(wp.stdinReader)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case value, ok := <-wp.cmd.Stdout:
				if !ok {
					return
				}
				wp.EmitEvent(ProcessLogEvent{Message: value})
			case value, ok := <-wp.cmd.Stderr:
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
		wp.gracefulStop(statusChan)
		err = wp.cmd.Stop()
	case <-signalChan:
		wp.gracefulStop(statusChan)
		err = wp.cmd.Stop()
	case status := <-statusChan:
		signal.Stop(signalChan)
		err = status.Error
	}
	wg.Wait()
	return err
}

func (wp *wrapper) gracefulStop(statusChan <-chan cmd.Status) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wp.gracefulTimeout)
	defer cancel()

	err := wp.gracefulStopper(wp)
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

func (wp *wrapper) runGrpcServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			return handler(WithWrapper(ctx, wp), req)
		}),
		grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, &defaultGrpcWrappedStream{ss, WithWrapper(ss.Context(), wp)})
		}),
	)
	for _, service := range wp.services {
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
