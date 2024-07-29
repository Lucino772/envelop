package wrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
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

func WithHooks(hooks []HookConfig) WrapperOptFunc {
	return func(options *wrapperOptions) {
		senders := make([]func(context.Context, []byte) error, 0)
		for _, hookConf := range hooks {
			if hookConf.Type == "http" {
				if sender := getWebhookSender(hookConf.Options); sender != nil {
					senders = append(senders, sender)
				}
			}
		}

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
					wg, _ := errgroup.WithContext(ctx)
					for _, sender := range senders {
						sender := sender
						wg.Go(func() error {
							return sender(ctx, data)
						})
					}
					_ = wg.Wait()
				}
			}
			return nil
		})
	}
}

func getWebhookSender(options map[string]any) func(context.Context, []byte) error {
	url, ok := options["url"]
	if !ok {
		return nil
	}

	return func(parent context.Context, data []byte) error {
		ctx, cancel := context.WithTimeout(parent, 10*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "POST", url.(string), bytes.NewBuffer(data))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// TODO: Do we except a response ? If so, what's the shape ?
		return nil
	}
}
