package minecraft

import (
	"context"

	"github.com/Lucino772/envelop/internal/wrapper"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"github.com/Lucino772/envelop/pkg/rcon"
	"google.golang.org/grpc"
)

type minecraftRconService struct {
	pb.UnimplementedRconServer
	wrapper wrapper.Wrapper
}

func NewMinecraftRconService(w wrapper.Wrapper) *minecraftRconService {
	return &minecraftRconService{wrapper: w}
}

func (service *minecraftRconService) SendCommand(ctx context.Context, req *pb.RconCommand) (*pb.RconResponse, error) {
	// TODO: Check if Rcon is enabled
	resp, err := rcon.Send(ctx, "localhost", 25575, "password", req.Value)
	if err != nil {
		return nil, err
	}
	return &pb.RconResponse{Value: resp}, nil
}

func (service *minecraftRconService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterRconServer(server, service)
}
