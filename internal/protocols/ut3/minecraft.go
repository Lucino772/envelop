package ut3

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"time"
)

type MinecraftQueryStats struct {
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

func QueryMinecraft(conn net.Conn, stats *MinecraftQueryStats) error {
	var sessionId = uint32(time.Now().Unix()) & 0x0F0F0F0F
	token, err := getChallengeToken(conn, int32(sessionId))
	if err != nil {
		return err
	}
	if err := sendQuery(conn, int32(sessionId), token); err != nil {
		return err
	}

	var buf bufio.Reader
	buf.Reset(conn)
	if _, err := buf.Discard(11); err != nil {
		return err
	}

	for {
		key, err := buf.ReadString(0x00)
		if err != nil {
			return err
		}
		key = key[:len(key)-1]
		if len(key) == 0 {
			break
		}

		value, err := buf.ReadString(0x00)
		if err != nil {
			return err
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
				return err
			}
			stats.Port = uint16(port)
		case "plugins":
			// TODO: Parse plugins
			stats.Plugins = []string{}
		case "numplayers":
			numPlayers, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			stats.NumPlayers = uint16(numPlayers)
		case "maxplayers":
			maxPlayers, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			stats.MaxPlayers = uint16(maxPlayers)
		}
	}

	if _, err := buf.Discard(10); err != nil {
		return err
	}

	stats.Players = make([]string, 0, stats.NumPlayers)
	for {
		player, err := buf.ReadString(0x00)
		if err != nil {
			return err
		}
		if len(player) == 1 {
			break
		}
		stats.Players = append(stats.Players, player[:len(player)-1])
	}
	return nil
}
