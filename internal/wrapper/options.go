package wrapper

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
)

type WrapperStopper func(WrapperContext) error
type WrapperOptFunc func(*wrapperOptions)

type WrapperTask func(context.Context) error

type WrapperService interface {
	Register(grpc.ServiceRegistrar)
}

type WrapperRegistrar interface {
	AddService(WrapperService)
	AddTask(WrapperTask)
}

type WrapperModule interface {
	Register(WrapperRegistrar)
}

type wrapperOptions struct {
	dir             string
	env             []string
	gracefulTimeout time.Duration
	gracefulStopper WrapperStopper
	services        []WrapperService
	tasks           []WrapperTask
}

func defaultOptions() wrapperOptions {
	return wrapperOptions{
		gracefulTimeout: 30 * time.Second,
		services:        make([]WrapperService, 0),
		tasks:           make([]WrapperTask, 0),
	}
}

func (options *wrapperOptions) AddService(service WrapperService) {
	options.services = append(options.services, service)
}

func (options *wrapperOptions) AddTask(task WrapperTask) {
	options.tasks = append(options.tasks, task)
}

func WithForwardLogToStdout() WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.AddTask(func(ctx context.Context) error {
			wp, err := FromContext(ctx)
			if err != nil {
				return err
			}

			sub := wp.SubscribeLogs()
			defer sub.Close()

			for item := range sub.Receive() {
				log.Println(item)
			}
			return nil
		})
	}
}

func WithWorkingDirectory(dir string) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.dir = dir
	}
}

func WithEnv(env []string) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.env = append(options.env, env...)
	}
}

func WithGracefulTimeout(timeout time.Duration) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.gracefulTimeout = timeout
	}
}

func WithGracefulStopper(stopper WrapperStopper) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.gracefulStopper = stopper
	}
}

func WithModule(module WrapperModule) WrapperOptFunc {
	return func(options *wrapperOptions) {
		module.Register(options)
	}
}

func WithHook(hook WrapperHook) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.AddTask(func(ctx context.Context) error {
			wp, err := FromContext(ctx)
			if err != nil {
				return err
			}

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
}
