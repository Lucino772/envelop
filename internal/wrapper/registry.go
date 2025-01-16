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

type Hook interface {
	Execute(context.Context, []byte) error
}

type Stopper func(Wrapper) error

type Task interface {
	Name() string
	Run(context.Context, Wrapper) error
}

type Service interface {
	Register(grpc.ServiceRegistrar)
}

type Registry struct {
	Tasks    []Task
	Services []Service

	Stoppers        map[string]func(map[string]any) Stopper
	LoggingHandlers map[string]func(map[string]any) slog.Handler
	Hooks           map[string]func(map[string]any) Hook
}

func NewRegistry() *Registry {
	return &Registry{
		Tasks:           make([]Task, 0),
		Services:        make([]Service, 0),
		Stoppers:        make(map[string]func(map[string]any) Stopper),
		LoggingHandlers: make(map[string]func(map[string]any) slog.Handler),
		Hooks:           make(map[string]func(map[string]any) Hook),
	}
}
