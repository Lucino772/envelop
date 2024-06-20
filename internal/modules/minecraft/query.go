package minecraft

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"strconv"
)

type QueryStats struct {
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

type buffer struct {
	*bufio.Reader
}

type queryClient struct {
	net.Conn

	sessId         int32
	challengeToken int32
	reader         *buffer
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

func (c *queryClient) QueryStats() (*QueryStats, error) {
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

	var stats QueryStats
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
