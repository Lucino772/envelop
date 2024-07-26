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

func (wp *Wrapper) runGrpcServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		wp.withUnaryInterceptor(),
		wp.withStreamInterceptor(),
	)
	for _, service := range wp.options.services {
		service.Register(grpcServer)
	}

	quit := make(chan error)
	go func() {
		defer close(quit)
		err := grpcServer.Serve(lis)
		if err != nil {
			quit <- err
		}
	}()

	select {
	case <-ctx.Done():
		grpcServer.Stop()
		return nil
	case err := <-quit:
		return err
	}
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
