package core

import (
	"errors"

	"github.com/Lucino772/envelop/internal/wrapper"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type coreSystemService struct {
	pb.UnimplementedSystemServer
}

func NewCoreSystemService() *coreSystemService {
	return &coreSystemService{}
}

func (service *coreSystemService) StreamEvents(_ *emptypb.Empty, stream pb.System_StreamEventsServer) error {
	wp, ok := wrapper.FromIncomingGrpcContext(stream.Context())
	if !ok {
		return errors.New("wrapper is not in context")
	}
	producer := wp.GetEventsProducer()
	channel := producer.Subscribe()
	defer producer.Unsubscribe(channel)

	for event := range channel {
		if err := stream.Send(event); err != nil {
			return err
		}
	}
	return nil
}

func (service *coreSystemService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterSystemServer(server, service)
}
