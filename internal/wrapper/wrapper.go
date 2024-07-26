package wrapper

import (
	"context"
	"io"
	"os"

	"github.com/Lucino772/envelop/pkg/pubsub"
	"github.com/go-cmd/cmd"
	"golang.org/x/sync/errgroup"
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
