package minecraft

import (
	"context"

	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"google.golang.org/grpc"
)

type minecraftRconService struct {
	pb.UnimplementedRconServer
}

func NewMinecraftRconService() *minecraftRconService {
	return &minecraftRconService{}
}

func (service *minecraftRconService) SendCommand(ctx context.Context, req *pb.RconCommand) (*pb.RconResponse, error) {
	// TODO: Check if Rcon is enabled
	resp, err := RconSend(ctx, "localhost", 25575, "password", req.Value)
	if err != nil {
		return nil, err
	}
	return &pb.RconResponse{Value: resp}, nil
}

func (service *minecraftRconService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterRconServer(server, service)
}
