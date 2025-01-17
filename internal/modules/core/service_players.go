package core

import (
	"context"

	"github.com/Lucino772/envelop/internal/wrapper"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type corePlayersService struct {
	pb.UnimplementedPlayersServer
}

func NewCorePlayersService() *corePlayersService {
	return &corePlayersService{}
}

func (service *corePlayersService) ListPlayers(ctx context.Context, _ *emptypb.Empty) (*pb.PlayerList, error) {
	wp, err := wrapper.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var state = wp.GetState()
	players := make([]*pb.Player, 0)
	for _, player := range state.Players.List {
		players = append(players, &pb.Player{Name: player.Id})
	}
	return &pb.PlayerList{
		NumPlayers: uint32(state.Players.Count),
		MaxPlayers: uint32(state.Players.Max),
		Players:    players,
	}, nil
}

func (service *corePlayersService) StreamPlayers(_ *emptypb.Empty, stream pb.Players_StreamPlayersServer) error {
	wp, err := wrapper.FromContext(stream.Context())
	if err != nil {
		return err
	}

	sub := wp.SubscribeStates()
	defer sub.Close()

	for state := range sub.Receive() {
		playerList := &pb.PlayerList{
			NumPlayers: uint32(state.Players.Count),
			MaxPlayers: uint32(state.Players.Max),
			Players:    make([]*pb.Player, 0),
		}
		for _, player := range state.Players.List {
			playerList.Players = append(playerList.Players, &pb.Player{Name: player.Id})
		}
		if err := stream.Send(playerList); err != nil {
			return err
		}
	}
	return nil
}

func (service *corePlayersService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterPlayersServer(server, service)
}
