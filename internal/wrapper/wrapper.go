package wrapper

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Lucino772/envelop/internal/utils"
	"github.com/go-cmd/cmd"
	"google.golang.org/grpc"
)

type Wrapper struct {
	ProcessStatusState WrapperStateAccessor[ProcessStatusState]
	PlayerState        WrapperStateAccessor[PlayerState]

	options        wrapperOptions
	cmd            *cmd.Cmd
	stdinReader    io.Reader
	stdinWriter    io.WriteCloser
	logsProducer   *utils.Producer[string]
	eventsProducer *utils.Producer[Event]
}

type wrapperStateProperty[T WrapperState] struct {
	stateObj T
	mu       sync.Mutex
	notify   func(WrapperState)
}

func newWrapperStateProperty[T WrapperState](initialState T, notify func(WrapperState)) *wrapperStateProperty[T] {
	return &wrapperStateProperty[T]{
		stateObj: initialState,
		notify:   notify,
	}
}

func (property *wrapperStateProperty[T]) Get() T {
	return property.stateObj
}

func (property *wrapperStateProperty[T]) Set(state T) {
	property.mu.Lock()
	defer property.mu.Unlock()
	if !property.stateObj.Equals(state) {
		property.stateObj = state
		property.notify(state)
	}
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
	command.Env = options.env

	wrapper := &Wrapper{
		options:        options,
		cmd:            command,
		stdinReader:    stdinReader,
		stdinWriter:    stdinWriter,
		logsProducer:   utils.NewProducer[string](),
		eventsProducer: utils.NewProducer[Event](),
	}
	wrapper.ProcessStatusState = newWrapperStateProperty(ProcessStatusState{
		Description: "Unknown",
	}, wrapper.updateState)
	wrapper.PlayerState = newWrapperStateProperty(PlayerState{
		Count:   0,
		Max:     0,
		Players: make([]string, 0),
	}, wrapper.updateState)
	return wrapper, nil
}

func (wp *Wrapper) WriteCommand(command string) error {
	_, err := wp.stdinWriter.Write([]byte(fmt.Sprintf("%s\n", command)))
	return err
}

func (wp *Wrapper) SubscribeLogs() WrapperSubscriber[string] {
	return wp.logsProducer.Subscribe()
}

func (wp *Wrapper) SubscribeEvents() WrapperSubscriber[Event] {
	return wp.eventsProducer.Subscribe()
}

func (wp *Wrapper) PublishEvent(event WrapperEvent) {
	wp.eventsProducer.Publish(Event{
		Id:        "", // TODO: Get Unique ID
		Timestamp: time.Now().Unix(),
		Name:      event.GetEventName(),
		Data:      event,
	})
}

func (wp *Wrapper) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(NewIncomingContext(parent, wp))
	defer cancel()

	wp.options.tasks = append(
		wp.options.tasks,
		wp.eventsProducer.Run,
		wp.logsProducer.Run,
	)
	grpcServer, err := wp.startGrpc()
	if err != nil {
		return err
	}
	defer grpcServer.Stop()
	for _, task := range wp.options.tasks {
		go task(ctx)
	}
	wp.runProcess(ctx)
	return nil
}

func (wp *Wrapper) startGrpc() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(getGrpcUnaryInterceptor(wp)),
		grpc.StreamInterceptor(getGrpcStreamInterceptor(wp)),
	)
	for _, service := range wp.options.services {
		service.Register(grpcServer)
	}
	go grpcServer.Serve(lis)
	return grpcServer, nil
}

func (wp *Wrapper) runProcess(ctx context.Context) {
	defer wp.stdinWriter.Close()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	statusChan := wp.cmd.StartWithStdin(wp.stdinReader)
	go wp.produceLogs()

	select {
	case <-ctx.Done():
		wp.gracefulStop(statusChan)
		wp.cmd.Stop()
	case <-signalChan:
		wp.gracefulStop(statusChan)
		wp.cmd.Stop()
	case <-statusChan:
		signal.Stop(signalChan)
	}
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

func (wp *Wrapper) produceLogs() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for value := range wp.cmd.Stdout {
			wp.logsProducer.Publish(value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range wp.cmd.Stderr {
			wp.logsProducer.Publish(value)
		}
	}()
	wg.Wait()
}

func (wp *Wrapper) updateState(state WrapperState) {
	wp.PublishEvent(StateUpdateEvent{
		Name: state.GetStateName(),
		Data: state,
	})
}
