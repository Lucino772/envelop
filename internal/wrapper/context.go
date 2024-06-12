package wrapper

import "context"

type wrapperIncomingGrpcKey struct{}

func FromIncomingContext(ctx context.Context) (*Wrapper, bool) {
	wrapper, ok := ctx.Value(wrapperIncomingGrpcKey{}).(*Wrapper)
	if !ok {
		return nil, false
	}
	return wrapper, true
}

func NewIncomingContext(ctx context.Context, wp *Wrapper) context.Context {
	return context.WithValue(ctx, wrapperIncomingGrpcKey{}, wp)
}
