package rcon

import (
	"context"

	"github.com/Lucino772/envelop/internal/protocols"
	"github.com/Lucino772/envelop/internal/wrapper"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"google.golang.org/grpc"
)

type rconService struct {
	pb.UnimplementedRconServer

	passwordKey string
	portKey     string
	enabledKey  string
}

func newRconService(passwordKey string, portKey string, enabledKey string) *rconService {
	return &rconService{
		passwordKey: passwordKey,
		portKey:     portKey,
		enabledKey:  enabledKey,
	}
}

func (service *rconService) SendCommand(ctx context.Context, req *pb.RconCommand) (*pb.RconResponse, error) {
	wp, err := wrapper.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	config, err := wp.GetServerConfig()
	if err != nil {
		return nil, err
	}

	if service.enabledKey != "" && !config.GetBool(service.enabledKey, false) {
		// TODO: Return error, rcon disabled
		return &pb.RconResponse{Value: ""}, nil
	}
	rconPort := config.GetInt16(service.portKey, -1)
	rconPassword := config.GetString(service.passwordKey, "")
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

func (service *rconService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterRconServer(server, service)
}
