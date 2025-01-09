package wrapper

import (
	"syscall"
)

func NewGracefulStopper(name string, options map[string]any) Stopper {
	switch name {
	case "cmd":
		command := options["cmd"].(string)
		return func(w Wrapper) error {
			return w.WriteStdin(command)
		}
	case "signal":
		sig := syscall.Signal(options["signal"].(int))
		return func(w Wrapper) error {
			return w.SendSignal(sig)
		}
	default:
		return nil
	}
}
