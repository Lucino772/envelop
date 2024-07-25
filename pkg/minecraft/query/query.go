package query

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"
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

func GetStats(conn net.Conn) (*QueryStats, error) {
	client := client{inner: conn}
	binary.BigEndian.PutUint32(
		client.sessionId[:],
		uint32(time.Now().Unix())&0x0F0F0F0F,
	)
	token, err := client.getChallengeToken()
	if err != nil {
		return nil, err
	}
	return client.getStats(token)
}

type client struct {
	inner     net.Conn
	sessionId [4]byte
}

func (client *client) getChallengeToken() ([]byte, error) {
	req := [7]byte{0xFE, 0xFD, 0x09}
	copy(req[3:], client.sessionId[:])
	if _, err := client.inner.Write(req[:]); err != nil {
		return nil, err
	}
	if err := client.checkRead(9); err != nil {
		return nil, err
	}

	buf := bufio.NewReader(client.inner)
	tokenString, err := buf.ReadString(0x00)
	if err != nil {
		return nil, err
	}
	token, err := strconv.ParseInt(tokenString[:len(tokenString)-1], 10, 32)
	if err != nil {
		return nil, err
	}
	var tokenBytes [4]byte
	binary.BigEndian.PutUint32(tokenBytes[:], uint32(token))
	return tokenBytes[:], nil
}

func (client *client) getStats(token []byte) (*QueryStats, error) {
	req := [15]byte{0xFE, 0xFD, 0x00}
	copy(req[3:], client.sessionId[0:])
	copy(req[7:], token)
	if _, err := client.inner.Write(req[:]); err != nil {
		return nil, err
	}
	if err := client.checkRead(0); err != nil {
		return nil, err
	}

	var stats QueryStats
	buf := bufio.NewReader(client.inner)
	if _, err := buf.Discard(11); err != nil {
		return nil, err
	}

	for i := 0; i < 10; i++ {
		key, err := buf.ReadString(0x00)
		if err != nil {
			return nil, err
		}
		key = key[:len(key)-1]

		value, err := buf.ReadString(0x00)
		if err != nil {
			return nil, err
		}
		value = value[:len(value)-1]

		switch strings.ToLower(key) {
		case "hostname":
			stats.Motd = value
		case "gametype":
			stats.GameType = value
		case "game_id":
			stats.GameId = value
		case "version":
			stats.Version = value
		case "map":
			stats.MapName = value
		case "hostip":
			stats.Host = value
		case "hostport":
			port, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			stats.Port = uint16(port)
		case "plugins":
			// TODO: Parse plugins
			stats.Plugins = []string{}
		case "numplayers":
			numPlayers, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			stats.NumPlayers = uint16(numPlayers)
		case "maxplayers":
			maxPlayers, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			stats.MaxPlayers = uint16(maxPlayers)
		}
	}

	if _, err := buf.Discard(11); err != nil {
		return nil, err
	}

	stats.Players = make([]string, 0)
	for {
		player, err := buf.ReadString(0x00)
		if err != nil {
			return nil, err
		}
		if len(player) == 1 {
			break
		}
		stats.Players = append(stats.Players, player[:len(player)-1])
	}

	return &stats, nil
}

func (client *client) checkRead(typ byte) error {
	var header [5]byte
	if _, err := client.inner.Read(header[0:]); err != nil {
		return err
	}
	if header[0] != typ && !bytes.Equal(header[1:], client.sessionId[:]) {
		return errors.New("invalid header")
	}
	return nil
}
