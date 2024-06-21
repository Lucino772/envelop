package wrapper

import "context"

type wrapperIncomingGrpcKey struct{}

func FromIncomingContext(ctx context.Context) (WrapperContext, bool) {
	wrapper, ok := ctx.Value(wrapperIncomingGrpcKey{}).(WrapperContext)
	if !ok {
		return nil, false
	}
	return wrapper, true
}

func NewIncomingContext(ctx context.Context, wp WrapperContext) context.Context {
	return context.WithValue(ctx, wrapperIncomingGrpcKey{}, wp)
}
