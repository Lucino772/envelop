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
