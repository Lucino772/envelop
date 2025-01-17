package minecraft

import (
	"context"
	"time"

	"github.com/Lucino772/envelop/internal/wrapper"
)

type fetchMinecraftPlayersTask struct{}

func NewFetchMinecraftPlayersTask() *fetchMinecraftPlayersTask {
	return &fetchMinecraftPlayersTask{}
}

func (task *fetchMinecraftPlayersTask) Name() string {
	return "watch-minecraft-players"
}

func (task *fetchMinecraftPlayersTask) Run(ctx context.Context, wp wrapper.Wrapper) error {
	err := task.waitServerReady(ctx, wp)
	if err != nil {
		return err
	}
	for {
		// TODO: Add config options for port and version
		stats, err := Query(ctx, "localhost", 25565, SLP_Version17)
		if err != nil {
			return err
		}
		wp.UpdateState(func(state wrapper.ServerState) wrapper.ServerState {
			state.Players.Count = stats.Players.Online
			state.Players.Max = stats.Players.Max
			state.Players.List = make([]wrapper.ServerState_Player, 0)
			for _, player := range stats.Players.Sample {
				state.Players.List = append(
					state.Players.List,
					wrapper.ServerState_Player{
						Id: player.Id,
						Attributes: map[string]any{
							"name": player.Name,
						},
					},
				)
			}
			return state
		})

		select {
		case <-time.After(time.Second * 3):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (task *fetchMinecraftPlayersTask) waitServerReady(ctx context.Context, wp wrapper.Wrapper) error {
	sub := wp.SubscribeLogs()
	defer sub.Close()
	messages := sub.Receive()

	for {
		select {
		case value := <-messages:
			if matches := serverReadyRegex.FindStringSubmatch(value); matches != nil {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
