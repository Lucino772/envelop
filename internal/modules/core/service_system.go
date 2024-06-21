package core

import (
	"encoding/json"

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
	wp, err := wrapper.FromContext(stream.Context())
	if err != nil {
		return err
	}
	sub := wp.SubscribeEvents()
	defer sub.Unsubscribe()

	for event := range sub.Messages() {
		evData, err := json.Marshal(event.Data)
		if err != nil {
			return err
		}
		grpcEvent := pb.Event{
			Id:   event.Id,
			Name: event.Name,
			Data: evData,
		}
		if err := stream.Send(&grpcEvent); err != nil {
			return err
		}
	}
	return nil
}

func (service *coreSystemService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterSystemServer(server, service)
}
