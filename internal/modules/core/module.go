package core

import (
	"context"
	"syscall"

	"github.com/Lucino772/envelop/internal/wrapper"
)

func Initialize(_ map[string]any, registry *wrapper.Registry) {
	registry.Services = append(
		registry.Services,
		NewCoreSystemService(),
		NewCoreProcessService(),
		NewCorePlayersService(),
	)
	registry.Tasks = append(
		registry.Tasks,
		wrapper.NewNamedTask(
			"process-logs-forward",
			func(ctx context.Context, wp wrapper.Wrapper) error {
				sub := wp.SubscribeLogs()
				defer sub.Close()

				logger := wp.Logger()
				for log := range sub.Receive() {
					logger.LogAttrs(ctx, LevelProcess, log)
				}
				return nil
			},
		),
	)
	registry.Stoppers["cmd"] = func(opts map[string]any) wrapper.Stopper {
		command := opts["cmd"].(string)
		return func(w wrapper.Wrapper) error {
			return w.WriteStdin(command)
		}
	}
	registry.Stoppers["signal"] = func(opts map[string]any) wrapper.Stopper {
		sig := syscall.Signal(opts["signal"].(int))
		return func(w wrapper.Wrapper) error {
			return w.SendSignal(sig)
		}
	}
	registry.LoggingHandlers["default"] = NewDefaultLoggingHandler
	registry.LoggingHandlers["http"] = NewHttpLoggingHandler
	registry.Hooks["http"] = NewHttpHook
}
