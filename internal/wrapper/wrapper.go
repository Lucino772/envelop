package wrapper

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/Lucino772/envelop/pkg/pubsub"
	"github.com/go-cmd/cmd"
	"golang.org/x/sync/errgroup"
)

type Wrapper struct {
	options        wrapperOptions
	cmd            *cmd.Cmd
	stdinReader    io.Reader
	stdinWriter    io.WriteCloser
	logsProducer   *pubsub.Producer[string]
	eventsProducer *pubsub.Producer[Event]

	processStatusState *stateProperty[ProcessStatusState]
	playerState        *stateProperty[PlayerState]
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
		logsProducer:   pubsub.NewProducer[string](5),
		eventsProducer: pubsub.NewProducer[Event](5),
	}
	wrapper.processStatusState = &stateProperty[ProcessStatusState]{
		state: ProcessStatusState{
			Description: "Unknown",
		},
		notify: wrapper.updateState,
	}
	wrapper.playerState = &stateProperty[PlayerState]{
		state: PlayerState{
			Count:   0,
			Max:     0,
			Players: []string{},
		},
		notify: wrapper.updateState,
	}
	return wrapper, nil
}

func (wp *Wrapper) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(wp.withContext(parent))
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

	errg, _ := errgroup.WithContext(ctx)
	errg.SetLimit(-1)
	for _, task := range wp.options.tasks {
		task := task
		errg.Go(func() error {
			return task(ctx)
		})
	}

	wp.runProcess(ctx)
	cancel()
	return errg.Wait()
}

func (wp *Wrapper) updateState(state WrapperState) {
	wp.PublishEvent(StateUpdateEvent{
		Name: state.GetStateName(),
		Data: state,
	})
}

type stateProperty[T WrapperState] struct {
	state  T
	mu     sync.Mutex
	notify func(WrapperState)
}

func (property *stateProperty[T]) Get() T {
	return property.state
}

func (property *stateProperty[T]) Set(state T) {
	property.mu.Lock()
	defer property.mu.Unlock()

	if !property.state.Equals(state) {
		property.state = state
		property.notify(state)
	}
}
