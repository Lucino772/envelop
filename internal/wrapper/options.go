package wrapper

import (
	"context"
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"time"

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
}

type Stopper func(Wrapper) error
type Module = func(Wrapper) []OptFunc
type Task = func(context.Context, Wrapper) error
type Service interface {
	Register(grpc.ServiceRegistrar)
}

func WithTask(t Task) OptFunc {
	return func(w *wrapper) {
		w.tasks = append(w.tasks, t)
	}
}

func WithService(s Service) OptFunc {
	return func(w *wrapper) {
		w.services = append(w.services, s)
	}
}

func WithModule(module Module) OptFunc {
	return func(w *wrapper) {
		for _, opt := range module(w) {
			opt(w)
		}
	}
}

func WithForwardLogToStdout() OptFunc {
	return WithTask(func(ctx context.Context, wp Wrapper) error {
		sub := wp.SubscribeLogs()
		defer sub.Close()

		for item := range sub.Receive() {
			log.Println(item)
		}
		return nil
	})
}

func WithWorkingDirectory(dir string) OptFunc {
	return func(w *wrapper) {
		w.cmd.Dir = dir
	}
}

func WithEnv(env []string) OptFunc {
	return func(w *wrapper) {
		w.cmd.Env = append(w.cmd.Env, env...)
	}
}

func WithGracefulTimeout(timeout time.Duration) OptFunc {
	return func(w *wrapper) {
		w.gracefulTimeout = timeout
	}
}

func WithGracefulStopper(stopper Stopper) OptFunc {
	return func(w *wrapper) {
		w.gracefulStopper = stopper
	}
}

func WithHook(hook Hook) OptFunc {
	return WithTask(func(ctx context.Context, wp Wrapper) error {
		sub := wp.SubscribeEvents()
		defer sub.Close()

		for event := range sub.Receive() {
			data, err := json.Marshal(event)
			if err == nil {
				// TODO: Handle error, log maybe ?
				_ = hook.Execute(ctx, data)
			}
		}
		return nil
	})
}
