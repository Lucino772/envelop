package wrapper

import (
	"context"
	"log"
	"os"
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
			defer sub.Unsubscribe()

			for item := range sub.Receive() {
				log.Println(item)
			}
			return nil
		})
	}
}

func WithForwardLogToEvent() WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.AddTask(func(ctx context.Context) error {
			wp, err := FromContext(ctx)
			if err != nil {
				return err
			}

			sub := wp.SubscribeLogs()
			defer sub.Unsubscribe()

			for item := range sub.Receive() {
				wp.PublishEvent(ProcessLogEvent{
					Value: item,
				})
			}
			return nil
		})
	}
}

func WithForwardStateToEvent() WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.AddTask(func(ctx context.Context) error {
			wp, err := FromContext(ctx)
			if err != nil {
				return err
			}

			sub := wp.SubscribeStates()
			defer sub.Unsubscribe()

			for state := range sub.Receive() {
				wp.PublishEvent(StateUpdateEvent{
					Name: state.GetStateName(),
					Data: state,
				})
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

func WithGracefulStopSignal(signal os.Signal) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.gracefulStopper = func(wp WrapperContext) error {
			return wp.SendSignal(signal)
		}
	}
}

func WithGracefulStopCommand(command string) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.gracefulStopper = func(wp WrapperContext) error {
			return wp.WriteCommand(command)
		}
	}
}

func WithModule(module WrapperModule) WrapperOptFunc {
	return func(options *wrapperOptions) {
		module.Register(options)
	}
}
