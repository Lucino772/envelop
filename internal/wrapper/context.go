package wrapper

import "context"

type wrapperIncomingGrpcKey struct{}

func FromIncomingContext(ctx context.Context) (*DefaultWrapper, bool) {
	wrapper, ok := ctx.Value(wrapperIncomingGrpcKey{}).(*DefaultWrapper)
	if !ok {
		return nil, false
	}
	return wrapper, true
}

func NewIncomingContext(ctx context.Context, wp *DefaultWrapper) context.Context {
	return context.WithValue(ctx, wrapperIncomingGrpcKey{}, wp)
}
