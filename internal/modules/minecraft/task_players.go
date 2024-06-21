package minecraft

import (
	"bufio"
	"context"
	"errors"
	"net"
	"time"

	"github.com/Lucino772/envelop/internal/wrapper"
)

type fetchMinecraftPlayersTask struct{}

func NewFetchMinecraftPlayersTask() *fetchMinecraftPlayersTask {
	return &fetchMinecraftPlayersTask{}
}

func (task *fetchMinecraftPlayersTask) Run(ctx context.Context) error {
	wp, err := wrapper.FromContext(ctx)
	if err != nil {
		return err
	}

	ready := task.waitQueryReady(ctx, wp)
	if !ready {
		return errors.New("query is not enabled")
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			stats, _ := task.queryStats(ctx)
			if stats != nil {
				wp.PlayerState().Set(wrapper.PlayerState{
					Count:   int(stats.NumPlayers),
					Max:     int(stats.MaxPlayers),
					Players: stats.Players,
				})
			}
		}
	}
}

func (task *fetchMinecraftPlayersTask) waitQueryReady(ctx context.Context, wp wrapper.WrapperContext) bool {
	sub := wp.SubscribeLogs()
	defer sub.Unsubscribe()
	messages := sub.Messages()

	// Wait for server to be ready
	var serverReady = false
	for !serverReady {
		select {
		case value := <-messages:
			if matches := serverReadyRegex.FindStringSubmatch(value); matches != nil {
				serverReady = true
			}
		case <-ctx.Done():
			return false
		}
	}

	// Wait for Query be ready with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var queryReady = false
	for !queryReady {
		select {
		case value := <-messages:
			if matches := serverRegexQueryReady.FindStringSubmatch(value); matches != nil {
				queryReady = true
			}
		case <-timeoutCtx.Done():
			return false
		}
	}

	return queryReady
}

func (task *fetchMinecraftPlayersTask) queryStats(ctx context.Context) (*QueryStats, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "udp", "localhost:25565")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return nil, err
	}

	client := &queryClient{
		Conn:   conn,
		sessId: int32(time.Now().Unix()) & 0x0F0F0F0F,
		reader: &buffer{bufio.NewReaderSize(conn, 1472)},
	}

	if err := client.Challenge(); err != nil {
		return nil, err
	}
	return client.QueryStats()
}
