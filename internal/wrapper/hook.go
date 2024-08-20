package wrapper

import (
	"context"
)

type Hook interface {
	Execute(context.Context, []byte) error
}

func NewHook(typ string, options map[string]any) Hook {
	switch typ {
	case "http":
		return NewHttpHook(options)
	default:
		return nil
	}
}
