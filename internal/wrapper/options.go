package wrapper

import (
	"context"
	"io/fs"
	"log/slog"
	"os"

	"github.com/Lucino772/envelop/pkg/pubsub"
	"google.golang.org/grpc"
)

type Wrapper interface {
	Files() fs.FS
	WriteStdin(command string) error
	SendSignal(signal os.Signal) error
	SubscribeLogs() pubsub.Subscriber[string]
	SubscribeEvents() pubsub.Subscriber[Event]
	EmitEvent(event any)
	ReadState(state any) bool
	SubscribeStates() pubsub.Subscriber[any]
	UpdateState(state any)
	Logger() *slog.Logger
}

type Stopper func(Wrapper) error

type Task interface {
	Name() string
	Run(context.Context, Wrapper) error
}

type Service interface {
	Register(grpc.ServiceRegistrar)
}

func WithForwardProcessLogsToLogger(options *Options) {
	options.Tasks = append(
		options.Tasks,
		NewNamedTask(
			"process-logs-forward",
			func(ctx context.Context, wp Wrapper) error {
				sub := wp.SubscribeLogs()
				defer sub.Close()

				logger := wp.Logger()
				for log := range sub.Receive() {
					logger.LogAttrs(
						ctx,
						LevelProcess,
						log,
					)
				}
				return nil
			},
		),
	)
}
