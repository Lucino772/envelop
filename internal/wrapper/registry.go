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
	GetState() ServerState
	SubscribeStates() pubsub.Subscriber[ServerState]
	UpdateState(func(ServerState) ServerState)
	Logger() *slog.Logger
	GetServerConfig() (Map, error)
}

type Map interface {
	GetInt8(string, int8) int8
	GetUint8(string, uint8) uint8
	GetInt16(string, int16) int16
	GetUint16(string, uint16) uint16
	GetInt32(string, int32) int32
	GetUint32(string, uint32) uint32
	GetInt64(string, int64) int64
	GetUint64(string, uint64) uint64
	GetBool(string, bool) bool
	GetString(string, string) string
	GetMap(string, Map) Map
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

type ConfigParser interface {
	Parse(map[string]any, Wrapper) error
}

type Registry struct {
	Tasks    []Task
	Services []Service

	Stoppers        map[string]func(map[string]any) Stopper
	LoggingHandlers map[string]func(map[string]any) slog.Handler
	Hooks           map[string]func(map[string]any) Hook
	ConfigParser    map[string]func(map[string]any) ConfigParser
}

func NewRegistry() *Registry {
	return &Registry{
		Tasks:           make([]Task, 0),
		Services:        make([]Service, 0),
		Stoppers:        make(map[string]func(map[string]any) Stopper),
		LoggingHandlers: make(map[string]func(map[string]any) slog.Handler),
		Hooks:           make(map[string]func(map[string]any) Hook),
		ConfigParser:    make(map[string]func(map[string]any) ConfigParser),
	}
}
