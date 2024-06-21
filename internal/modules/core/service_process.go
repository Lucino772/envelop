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

func (service *coreProcessService) GetStatus(ctx context.Context, _ *emptypb.Empty) (*pb.Status, error) {
	wp, ok := wrapper.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("wrapper is not in context")
	}
	state := wp.ProcessStatusState().Get()
	return &pb.Status{
		Value: state.Description,
	}, nil
}

func (service *coreProcessService) StreamStatus(_ *emptypb.Empty, stream pb.Process_StreamStatusServer) error {
	wp, ok := wrapper.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("wrapper is not in context")
	}

	sub := wp.SubscribeEvents()
	defer sub.Unsubscribe()

	for event := range sub.Messages() {
		if event.Name == "/state/update" {
			if eventData, ok := event.Data.(wrapper.StateUpdateEvent); ok {
				if state, ok := eventData.Data.(wrapper.ProcessStatusState); ok {
					status := &pb.Status{
						Value: state.Description,
					}
					if err := stream.Send(status); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (service *coreProcessService) WriteCommand(ctx context.Context, request *pb.Command) (*emptypb.Empty, error) {
	wp, ok := wrapper.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("wrapper is not in context")
	}
	if err := wp.WriteCommand(request.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (service *coreProcessService) StreamLogs(_ *emptypb.Empty, stream pb.Process_StreamLogsServer) error {
	wp, ok := wrapper.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("wrapper is not in context")
	}
	sub := wp.SubscribeLogs()
	defer sub.Unsubscribe()

	for log := range sub.Messages() {
		if err := stream.Send(&pb.Log{Value: log}); err != nil {
			return err
		}
	}
	return nil
}

func (service *coreProcessService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterProcessServer(server, service)
}
