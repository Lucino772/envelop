package wrapper

import (
	"context"
	"log"
	"os"
	"time"
)

type WrapperStopper func(*Wrapper) error
type WrapperOptFunc func(*wrapperOptions)

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
		options.AddTask(func(ctx context.Context) {
			wp, ok := FromIncomingContext(ctx)
			if !ok {
				return
			}

			sub := wp.SubscribeLogs()
			defer sub.Unsubscribe()

			for item := range sub.Messages() {
				log.Println(item)
			}
		})
	}
}

func WithForwardLogToEvent() WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.AddTask(func(ctx context.Context) {
			wp, ok := FromIncomingContext(ctx)
			if !ok {
				return
			}

			sub := wp.SubscribeLogs()
			defer sub.Unsubscribe()

			for item := range sub.Messages() {
				wp.PublishEvent(ProcessLogEvent{
					Value: item,
				})
			}
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
		options.gracefulStopper = func(wp *Wrapper) error {
			process, err := os.FindProcess(wp.cmd.Status().PID)
			if err != nil {
				return err
			}
			process.Signal(signal)
			return nil
		}
	}
}

func WithGracefulStopCommand(command string) WrapperOptFunc {
	return func(options *wrapperOptions) {
		options.gracefulStopper = func(wp *Wrapper) error {
			return wp.WriteCommand(command)
		}
	}
}

func WithModule(module WrapperModule) WrapperOptFunc {
	return func(options *wrapperOptions) {
		module.Register(options)
	}
}
