package wrapper

import (
	"context"
	"io"
	"net"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"github.com/Lucino772/envelop/pkg/pubsub"
	"github.com/go-cmd/cmd"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Wrapper struct {
	options        wrapperOptions
	cmd            *cmd.Cmd
	stdinReader    io.Reader
	stdinWriter    io.WriteCloser
	logsProducer   pubsub.Publisher[string]
	eventsProducer pubsub.Publisher[Event]
	stateManager   *StatePublisher
}

type defaultGrpcWrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *defaultGrpcWrappedStream) Context() context.Context {
	return w.ctx
}

func NewWrapper(program string, args []string, opts ...WrapperOptFunc) (*Wrapper, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	options := defaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	command := cmd.NewCmdOptions(cmd.Options{
		Buffered:  false,
		Streaming: true,
	}, program, args...)
	command.Dir = options.dir
	command.Env = slices.Concat(os.Environ(), options.env)

	wrapper := &Wrapper{
		options:        options,
		cmd:            command,
		stdinReader:    stdinReader,
		stdinWriter:    stdinWriter,
		logsProducer:   pubsub.NewPublisher[string](5, nil),
		eventsProducer: pubsub.NewPublisher[Event](5, nil),
		stateManager:   NewStatePublisher(5),
	}
	return wrapper, nil
}

func (wp *Wrapper) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(wp.withContext(parent))
	defer cancel()

	wp.options.tasks = append(
		wp.options.tasks,
		wp.eventsProducer.Run,
		wp.stateManager.Run,
		wp.logsProducer.Run,
	)

	var (
		mainErrGroup, mainCtx = errgroup.WithContext(ctx)
		taskErrGroup, _       = errgroup.WithContext(mainCtx)
	)
	mainErrGroup.SetLimit(2)
	taskErrGroup.SetLimit(-1)

	for _, task := range wp.options.tasks {
		task := task
		taskErrGroup.Go(func() error {
			return task(mainCtx)
		})
	}

	mainErrGroup.Go(func() error {
		defer cancel()
		return wp.runGrpcServer(mainCtx)
	})
	mainErrGroup.Go(func() error {
		defer cancel()
		return wp.runProcess(mainCtx)
	})

	err := mainErrGroup.Wait()
	taskErrGroup.Wait()
	return err
}

func (wp *Wrapper) runProcess(ctx context.Context) error {
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
			case value, closed := <-wp.cmd.Stdout:
				if !closed {
					return
				}
				wp.logsProducer.Publish(value)
			case value, closed := <-wp.cmd.Stderr:
				if !closed {
					return
				}
				wp.logsProducer.Publish(value)
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

func (wp *Wrapper) gracefulStop(statusChan <-chan cmd.Status) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wp.options.gracefulTimeout)
	defer cancel()

	err := wp.options.gracefulStopper(wp)
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

func (wp *Wrapper) runGrpcServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				return handler(wp.withContext(ctx), req)
			},
		),
		grpc.StreamInterceptor(
			func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				return handler(srv, &defaultGrpcWrappedStream{ss, wp.withContext(ss.Context())})
			},
		),
	)
	for _, service := range wp.options.services {
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
