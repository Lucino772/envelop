package core

import (
	"context"

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

func (service *coreProcessService) GetStatus(ctx context.Context, _ *emptypb.Empty) (*pb.Status, error) {
	wp, err := wrapper.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	var state = wp.GetState()
	return &pb.Status{Value: state.Status.Description}, nil
}

func (service *coreProcessService) StreamStatus(_ *emptypb.Empty, stream pb.Process_StreamStatusServer) error {
	wp, err := wrapper.FromContext(stream.Context())
	if err != nil {
		return err
	}

	sub := wp.SubscribeStates()
	defer sub.Close()

	for state := range sub.Receive() {
		status := &pb.Status{Value: state.Status.Description}
		if err := stream.Send(status); err != nil {
			return err
		}
	}
	return nil
}

func (service *coreProcessService) WriteCommand(ctx context.Context, request *pb.Command) (*emptypb.Empty, error) {
	wp, err := wrapper.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	if err := wp.WriteStdin(request.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (service *coreProcessService) StreamLogs(_ *emptypb.Empty, stream pb.Process_StreamLogsServer) error {
	wp, err := wrapper.FromContext(stream.Context())
	if err != nil {
		return err
	}

	sub := wp.SubscribeLogs()
	defer sub.Close()

	for log := range sub.Receive() {
		if err := stream.Send(&pb.Log{Value: log}); err != nil {
			return err
		}
	}
	return nil
}

func (service *coreProcessService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterProcessServer(server, service)
}
