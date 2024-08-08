package wrapper

import (
	"syscall"
)

func NewGracefulStopper(name string, options map[string]any) WrapperStopper {
	switch name {
	case "cmd":
		command := options["cmd"].(string)
		return func(wp WrapperContext) error {
			return wp.WriteCommand(command)
		}
	case "signal":
		sig := options["signal"].(syscall.Signal)
		return func(wp WrapperContext) error {
			return wp.SendSignal(sig)
		}
	default:
		return nil
	}
}
