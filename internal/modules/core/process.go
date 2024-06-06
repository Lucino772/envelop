package core

import (
	"context"
	"errors"

	"github.com/Lucino772/envelop/internal/wrapper"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type coreProcessService struct {
	pb.UnimplementedProcessServer
}

func NewCoreProcessService() *coreProcessService {
	return &coreProcessService{}
}

func (service *coreProcessService) WriteCommand(ctx context.Context, request *pb.Command) (*emptypb.Empty, error) {
	wp, ok := wrapper.FromIncomingGrpcContext(ctx)
	if !ok {
		return nil, errors.New("wrapper is not in context")
	}
	if err := wp.WriteCommand(request.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (service *coreProcessService) StreamLogs(_ *emptypb.Empty, stream pb.Process_StreamLogsServer) error {
	wp, ok := wrapper.FromIncomingGrpcContext(stream.Context())
	if !ok {
		return errors.New("wrapper is not in context")
	}
	producer := wp.GetLogsProducer()
	channel := producer.Subscribe()
	defer producer.Unsubscribe(channel)

	for log := range channel {
		if err := stream.Send(&pb.Log{Value: log}); err != nil {
			return err
		}
	}
	return nil
}

func (service *coreProcessService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterProcessServer(server, service)
}
