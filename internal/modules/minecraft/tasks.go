package minecraft

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"

	"github.com/Lucino772/envelop/internal/wrapper"
)

var (
	serverStartingRegex   = regexp.MustCompile(`\[Server thread\/INFO\]\: Starting Minecraft server on \*:(?P<port>[0-9]+)`)
	serverPreparingRegex  = regexp.MustCompile(`\[(.*?)\]: Preparing spawn area: (?P<progress>[0-9]+)%`)
	serverReadyRegex      = regexp.MustCompile(`\[Server thread\/INFO\]\: Done \((.*?)\)! For help, type \"help\"`)
	serverStoppingRegex   = regexp.MustCompile(`\[Server thread\/INFO\]\: Stopping server`)
	serverRegexQueryReady = regexp.MustCompile(`\[Query Listener #1\/INFO\]\: Query running on (.*?):(?P<port>[0-9]+)`)
)

type buffer struct {
	*bufio.Reader
}

type queryStats struct {
	Motd       string
	GameType   string
	GameId     string
	Version    string
	MapName    string
	Host       string
	Port       uint16
	Plugins    []string
	NumPlayers uint16
	MaxPlayers uint16
	Players    []string
}

type queryClient struct {
	net.Conn

	sessId         int32
	challengeToken int32
	reader         *buffer
}

type checkMinecraftStatusTask struct{}

type fetchMinecraftPlayersTask struct{}

func NewCheckMinecraftStatusTask() *checkMinecraftStatusTask {
	return &checkMinecraftStatusTask{}
}

func (task *checkMinecraftStatusTask) processSubexpNames(regex *regexp.Regexp, matches []string) map[string]string {
	result := make(map[string]string)
	for ix, name := range regex.SubexpNames() {
		if ix != 0 && name != "" {
			result[name] = matches[ix]
		}
	}
	return result
}

func (task *checkMinecraftStatusTask) processValue(wp *wrapper.Wrapper, value string) {
	if matches := serverStartingRegex.FindStringSubmatch(value); matches != nil {
		wp.ProcessStatusState.Set(wrapper.ProcessStatusState{
			Description: "Starting",
		})
	} else if matches := serverPreparingRegex.FindStringSubmatch(value); matches != nil {
		groups := task.processSubexpNames(serverPreparingRegex, matches)
		wp.ProcessStatusState.Set(wrapper.ProcessStatusState{
			Description: fmt.Sprintf("Preparing (%s%%)", groups["progress"]),
		})
	} else if matches := serverReadyRegex.FindStringSubmatch(value); matches != nil {
		wp.ProcessStatusState.Set(wrapper.ProcessStatusState{
			Description: "Ready",
		})
	} else if matches := serverStoppingRegex.FindStringSubmatch(value); matches != nil {
		wp.ProcessStatusState.Set(wrapper.ProcessStatusState{
			Description: "Stopping",
		})
	}
}

func (task *checkMinecraftStatusTask) Run(ctx context.Context) {
	wp, ok := wrapper.FromIncomingContext(ctx)
	if !ok {
		return
	}

	sub := wp.SubscribeLogs()
	defer sub.Unsubscribe()
	messages := sub.Messages()

	for {
		select {
		case value := <-messages:
			task.processValue(wp, value)
		case <-ctx.Done():
			return
		}
	}
}

func NewFetchMinecraftPlayersTask() *fetchMinecraftPlayersTask {
	return &fetchMinecraftPlayersTask{}
}

func (task *fetchMinecraftPlayersTask) Run(ctx context.Context) {
	wp, ok := wrapper.FromIncomingContext(ctx)
	if !ok {
		return
	}

	ready := task.waitQueryReady(ctx, wp)
	if !ready {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			stats, _ := task.queryStats(ctx)
			if stats != nil {
				wp.PlayerState.Set(wrapper.PlayerState{
					Count:   int(stats.NumPlayers),
					Max:     int(stats.MaxPlayers),
					Players: stats.Players,
				})
			}
		}
	}
}

func (task *fetchMinecraftPlayersTask) waitQueryReady(ctx context.Context, wp *wrapper.Wrapper) bool {
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

func (task *fetchMinecraftPlayersTask) queryStats(ctx context.Context) (*queryStats, error) {
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

func (c *queryClient) Challenge() error {
	if _, err := c.writePacket(9); err != nil {
		return err
	}

	typ, sessId, err := c.readPacketHeader()
	if err != nil {
		return err
	}
	if typ != 9 || sessId != c.sessId {
		return errors.New("invalid packet header")
	}

	tokenString, err := c.reader.ReadNullTerminatedString()
	if err != nil {
		return err
	}
	token, err := strconv.Atoi(tokenString)
	if err != nil {
		return err
	}
	c.challengeToken = int32(token)
	return nil
}

func (c *queryClient) QueryStats() (*queryStats, error) {
	if _, err := c.writePacket(0, int32(c.challengeToken), uint32(0xFFFFFF01)); err != nil {
		return nil, err
	}

	typ, sessId, err := c.readPacketHeader()
	if err != nil {
		return nil, err
	}
	if typ != 0 || sessId != c.sessId {
		return nil, errors.New("invalid packet header")
	}

	var stats queryStats
	if _, err := c.reader.Discard(11); err != nil {
		return nil, err
	}

	info := make(map[string]string, 0)
	for i := 0; i < 10; i++ {
		key, err := c.reader.ReadNullTerminatedString()
		if err != nil {
			return nil, err
		}
		val, err := c.reader.ReadNullTerminatedString()
		if err != nil {
			return nil, err
		}
		info[key] = val
	}
	stats.Motd = info["hostname"]
	stats.GameType = info["gametype"]
	stats.GameId = info["game_id"]
	stats.Version = info["version"]
	stats.MapName = info["map"]
	stats.Host = info["hosip"]

	port, err := strconv.Atoi(info["hostport"])
	if err != nil {
		return nil, err
	}
	stats.Port = uint16(port)
	stats.Plugins = []string{}

	numPlayers, err := strconv.Atoi(info["numplayers"])
	if err != nil {
		return nil, err
	}
	stats.NumPlayers = uint16(numPlayers)
	maxPlayers, err := strconv.Atoi(info["maxplayers"])
	if err != nil {
		return nil, err
	}
	stats.MaxPlayers = uint16(maxPlayers)

	if _, err := c.reader.Discard(11); err != nil {
		return nil, err
	}

	stats.Players = make([]string, 0)

	player, err := c.reader.ReadNullTerminatedString()
	if err != nil {
		return nil, err
	}
	for len(player) != 0 {
		stats.Players = append(stats.Players, player)
		player, err = c.reader.ReadNullTerminatedString()
		if err != nil {
			return nil, err
		}
	}
	return &stats, nil
}

func (c *queryClient) writePacket(typ byte, data ...any) (int, error) {
	toWrite := []any{uint16(0xFEFD), byte(typ), int32(c.sessId)}
	toWrite = append(toWrite, data...)

	var buf bytes.Buffer
	for _, d := range toWrite {
		if err := binary.Write(&buf, binary.BigEndian, d); err != nil {
			return 0, err
		}
	}
	return c.Write(buf.Bytes())
}

func (c *queryClient) readPacketHeader() (typ byte, sessId int32, err error) {
	if err = binary.Read(c.reader, binary.BigEndian, &typ); err != nil {
		return typ, sessId, err
	}
	if err = binary.Read(c.reader, binary.BigEndian, &sessId); err != nil {
		return typ, sessId, err
	}
	return typ, sessId, err
}

func (b *buffer) ReadNullTerminatedString() (string, error) {
	data, err := b.ReadBytes(0)
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", nil
	}
	return string(data[0 : len(data)-1]), nil
}
