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
	state := wp.PlayerState().Get()
	players := make([]*pb.Player, 0)
	for _, player := range state.Players {
		players = append(players, &pb.Player{Name: player})
	}
	return &pb.PlayerList{
		NumPlayers: uint32(state.Count),
		MaxPlayers: uint32(state.Max),
		Players:    players,
	}, nil
}

func (service *corePlayersService) StreamPlayers(_ *emptypb.Empty, stream pb.Players_StreamPlayersServer) error {
	wp, err := wrapper.FromContext(stream.Context())
	if err != nil {
		return err
	}

	sub := wp.SubscribeEvents()
	defer sub.Unsubscribe()

	for event := range sub.Messages() {
		if event.Name == "/state/update" {
			if eventData, ok := event.Data.(wrapper.StateUpdateEvent); ok {
				if state, ok := eventData.Data.(wrapper.PlayerState); ok {
					playerList := &pb.PlayerList{
						NumPlayers: uint32(state.Count),
						MaxPlayers: uint32(state.Max),
						Players:    make([]*pb.Player, 0),
					}
					for _, player := range state.Players {
						playerList.Players = append(playerList.Players, &pb.Player{Name: player})
					}
					if err := stream.Send(playerList); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (service *corePlayersService) Register(server grpc.ServiceRegistrar) {
	pb.RegisterPlayersServer(server, service)
}
