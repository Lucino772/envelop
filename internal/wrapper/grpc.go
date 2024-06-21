package wrapper

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type defaultGrpcWrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *defaultGrpcWrappedStream) Context() context.Context {
	return w.ctx
}

func (wp *Wrapper) startGrpc() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer(
		wp.withUnaryInterceptor(),
		wp.withStreamInterceptor(),
	)
	for _, service := range wp.options.services {
		service.Register(grpcServer)
	}
	go grpcServer.Serve(lis)
	return grpcServer, nil
}

func (wp *Wrapper) withUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			return handler(wp.withContext(ctx), req)
		},
	)
}

func (wp *Wrapper) withStreamInterceptor() grpc.ServerOption {
	return grpc.StreamInterceptor(
		func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, &defaultGrpcWrappedStream{ss, wp.withContext(ss.Context())})
		},
	)
}
