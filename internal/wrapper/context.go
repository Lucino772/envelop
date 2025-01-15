package wrapper

import (
	"context"
	"errors"
)

var ErrWrapperContextMissing = errors.New("wrapper context missing")

type wrapperContextKey struct{}

func WithWrapper(ctx context.Context, wp Wrapper) context.Context {
	return context.WithValue(ctx, wrapperContextKey{}, wp)
}

func FromContext(ctx context.Context) (Wrapper, error) {
	value := ctx.Value(wrapperContextKey{})
	if value == nil {
		return nil, ErrWrapperContextMissing
	}
	wrapper, ok := value.(Wrapper)
	if !ok {
		return nil, ErrWrapperContextMissing
	}
	return wrapper, nil
}
