package minecraft

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/Lucino772/envelop/internal/wrapper"
	"github.com/Lucino772/envelop/pkg/minecraft/query"
	"golang.org/x/sync/errgroup"
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

	result := make(chan *query.QueryStats)
	defer close(result)

	var d net.Dialer
	wg := new(errgroup.Group)
	wg.SetLimit(1)
	for {
		conn, err := d.DialContext(ctx, "udp", "localhost:25565")
		if err != nil {
			return err
		}
		wg.Go(func() error {
			stats, err := query.GetStats(conn)
			if err != nil {
				return err
			}
			result <- stats
			return nil
		})
		select {
		case <-ctx.Done():
			conn.Close()
		case stats := <-result:
			if stats != nil {
				wp.PublishState(wrapper.PlayerState{
					Count:   int(stats.NumPlayers),
					Max:     int(stats.MaxPlayers),
					Players: stats.Players,
				})
			}
		}
		if err := wg.Wait(); err != nil {
			return err
		}
	}
}

func (task *fetchMinecraftPlayersTask) waitQueryReady(ctx context.Context, wp wrapper.WrapperContext) bool {
	sub := wp.SubscribeLogs()
	defer sub.Close()
	messages := sub.Receive()

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
