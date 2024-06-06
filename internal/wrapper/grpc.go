package wrapper

import (
	"context"

	"google.golang.org/grpc"
)

type defaultGrpcWrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func newDefaultGrpcWrappedStream(ctx context.Context, ss grpc.ServerStream) *defaultGrpcWrappedStream {
	return &defaultGrpcWrappedStream{ss, ctx}
}

func (w *defaultGrpcWrappedStream) Context() context.Context {
	return w.ctx
}

type wrapperIncomingGrpcKey struct{}

func FromIncomingGrpcContext(ctx context.Context) (*DefaultWrapper, bool) {
	wrapper, ok := ctx.Value(wrapperIncomingGrpcKey{}).(*DefaultWrapper)
	if !ok {
		return nil, false
	}
	return wrapper, true
}

func newIncomingGrpcContext(ctx context.Context, wp *DefaultWrapper) context.Context {
	return context.WithValue(ctx, wrapperIncomingGrpcKey{}, wp)
}

func getGrpcUnaryInterceptor(wp *DefaultWrapper) func(parent context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return handler(newIncomingGrpcContext(ctx, wp), req)
	}
}

func getGrpcStreamInterceptor(wp *DefaultWrapper) func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(
			srv, newDefaultGrpcWrappedStream(newIncomingGrpcContext(ss.Context(), wp), ss),
		)
	}
}
