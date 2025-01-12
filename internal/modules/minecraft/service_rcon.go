package minecraft

import (
	"context"
	"io/fs"

	"github.com/Lucino772/envelop/internal/protocols"
	"github.com/Lucino772/envelop/internal/wrapper"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"github.com/magiconair/properties"
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
	data, err := fs.ReadFile(service.wrapper.Files(), "server.properties")
	if err != nil {
		return nil, err
	}
	props, err := properties.LoadString(string(data))
	if err != nil {
		return nil, err
	}
	if !props.GetBool("enable-rcon", false) {
		// TODO: Return error, rcon disabled
		return &pb.RconResponse{Value: ""}, nil
	}
	rconPort := props.GetInt("rcon.port", -1)
	rconPassword := props.GetString("rcon.password", "")
	if rconPort == -1 || rconPassword == "" {
		// TODO: Return error, password or port not set
		return &pb.RconResponse{Value: ""}, nil
	}

	resp, err := protocols.SendRcon(ctx, "localhost", uint16(rconPort), rconPassword, req.Value)
	if err != nil {
		return nil, err
	}
	return &pb.RconResponse{Value: resp}, nil
}

func (service *minecraftRconService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterRconServer(server, service)
}
