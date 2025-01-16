package minecraft

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
	"unicode/utf16"
)

const (
	SLP_Version17   = 0
	SLP_Version16   = 1
	SLP_Version15   = 2
	SLP_VersionBeta = 3
)

type Stats struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"sample"`
	} `json:"players"`
	Description        map[string]any `json:"description"`
	Favicon            string         `json:"favicon"`
	EnforcesSecureChat bool           `json:"enforcesSecureChat"`
}

func Query(ctx context.Context, hostname string, port uint16, version int) (*Stats, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var pinger func(net.Conn, *Stats) error
	switch version {
	case SLP_Version17:
		pinger = func(c net.Conn, s *Stats) error {
			return ping17(c, hostname, port, s)
		}
	case SLP_Version16:
		pinger = func(c net.Conn, s *Stats) error {
			return ping16(c, hostname, uint32(port), s)
		}
	case SLP_Version15:
		pinger = ping15
	case SLP_VersionBeta:
		pinger = pingBeta
	}
	if pinger == nil {
		return nil, errors.New("invalid version")
	}

	var stats Stats
	errChan := make(chan error)
	defer close(errChan)

	go func() {
		errChan <- pinger(conn, &stats)
	}()

	select {
	case <-ctx.Done():
		conn.Close()
		<-errChan // wait for goroutine to exit
		return nil, ctx.Err()
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
		return &stats, nil
	}
}

func ping17(conn net.Conn, hostname string, port uint16, stats *Stats) error {
	handshake := makePacket17(0x00, makeHandshake17([]byte(hostname), port))
	if _, err := conn.Write(handshake); err != nil {
		return err
	}

	statusRequest := makePacket17(0x00, []byte{})
	if _, err := conn.Write(statusRequest); err != nil {
		return err
	}

	conn.SetDeadline(time.Now().Add(time.Second * 15))
	buff := bufio.NewReader(conn)
	packetLen, err := binary.ReadUvarint(buff)
	if err != nil {
		return err
	}
	packetData := make([]byte, packetLen)
	if _, err := io.ReadFull(buff, packetData); err != nil {
		return err
	}

	packetBuff := bufio.NewReader(bytes.NewReader(packetData))
	_, err = binary.ReadUvarint(packetBuff)
	if err != nil {
		return err
	}
	respLen, err := binary.ReadUvarint(packetBuff)
	if err != nil {
		return err
	}
	jsonData := make([]byte, respLen)
	if _, err := io.ReadFull(packetBuff, jsonData); err != nil {
		return err
	}
	return json.Unmarshal(jsonData, stats)
}

func ping16(conn net.Conn, hostname string, port uint32, stats *Stats) error {
	request := makePacket16(hostname, port)
	if _, err := conn.Write(request); err != nil {
		return err
	}

	conn.SetDeadline(time.Now().Add(time.Second * 15))
	buff := bufio.NewReader(conn)
	_, err := buff.ReadByte()
	if err != nil {
		return err
	}

	var packetLen int16
	if err := binary.Read(buff, binary.BigEndian, &packetLen); err != nil {
		return err
	}

	dataBuff := make([]uint16, packetLen)
	for i := 0; i < int(packetLen); i++ {
		if err := binary.Read(buff, binary.BigEndian, &dataBuff[i]); err != nil {
			return err
		}
	}
	decoded := []byte(string(utf16.Decode(dataBuff)))
	results := bytes.Split(decoded, []byte{0x00})

	protocol, err := strconv.Atoi(string(results[1]))
	if err != nil {
		return err
	}
	stats.Version.Protocol = protocol
	stats.Version.Name = string(results[2])
	stats.Description = map[string]any{
		"text": string(results[3]),
	}

	currentPlayer, err := strconv.Atoi(string(results[4]))
	if err != nil {
		return err
	}
	stats.Players.Online = currentPlayer
	maxPlayers, err := strconv.Atoi(string(results[5]))
	if err != nil {
		return err
	}
	stats.Players.Max = maxPlayers
	return nil
}

func ping15(conn net.Conn, stats *Stats) error {
	if _, err := conn.Write([]byte{0xFE, 0x01}); err != nil {
		return err
	}

	conn.SetDeadline(time.Now().Add(time.Second * 15))
	buff := bufio.NewReader(conn)
	_, err := buff.ReadByte()
	if err != nil {
		return err
	}

	var packetLen int16
	if err := binary.Read(buff, binary.BigEndian, &packetLen); err != nil {
		return err
	}

	dataBuff := make([]uint16, packetLen)
	for i := 0; i < int(packetLen); i++ {
		if err := binary.Read(buff, binary.BigEndian, &dataBuff[i]); err != nil {
			return err
		}
	}
	decoded := []byte(string(utf16.Decode(dataBuff)))
	results := bytes.Split(decoded, []byte{0x00})

	protocol, err := strconv.Atoi(string(results[1]))
	if err != nil {
		return err
	}
	stats.Version.Protocol = protocol
	stats.Version.Name = string(results[2])
	stats.Description = map[string]any{
		"text": string(results[3]),
	}

	currentPlayer, err := strconv.Atoi(string(results[4]))
	if err != nil {
		return err
	}
	stats.Players.Online = currentPlayer
	maxPlayers, err := strconv.Atoi(string(results[5]))
	if err != nil {
		return err
	}
	stats.Players.Max = maxPlayers
	return nil
}

func pingBeta(conn net.Conn, stats *Stats) error {
	if _, err := conn.Write([]byte{0xFE}); err != nil {
		return err
	}

	conn.SetDeadline(time.Now().Add(time.Second * 15))
	buff := bufio.NewReader(conn)
	_, _ = buff.ReadByte()

	var packetLen int16
	if err := binary.Read(buff, binary.BigEndian, &packetLen); err != nil {
		return err
	}

	dataBuff := make([]uint16, packetLen)
	for i := 0; i < int(packetLen); i++ {
		if err := binary.Read(buff, binary.BigEndian, &dataBuff[i]); err != nil {
			return err
		}
	}
	decoded := []byte(string(utf16.Decode(dataBuff)))
	results := bytes.Split(decoded, []byte{194, 167})

	stats.Version.Name = "beta"
	stats.Description = map[string]any{
		"text": string(results[0]),
	}
	currentPlayer, err := strconv.Atoi(string(results[1]))
	if err != nil {
		return err
	}
	stats.Players.Online = currentPlayer
	maxPlayers, err := strconv.Atoi(string(results[2]))
	if err != nil {
		return err
	}
	stats.Players.Max = maxPlayers
	return nil
}

func makePacket17(packetId int, data []byte) []byte {
	payload := make([]byte, 0)
	payload = binary.AppendUvarint(payload, uint64(packetId))
	payload = append(payload, data...)

	buff := make([]byte, 0)
	buff = binary.AppendUvarint(buff, uint64(len(payload)))
	buff = append(buff, payload...)
	return buff
}

func makeHandshake17(hostname []byte, port uint16) []byte {
	buff := make([]byte, 0)
	buff = binary.AppendUvarint(buff, 0)
	buff = binary.AppendUvarint(buff, uint64(len(hostname)))
	buff = append(buff, hostname...)
	buff = binary.BigEndian.AppendUint16(buff, port)
	buff = binary.AppendUvarint(buff, 1)
	return buff
}

func makePacket16(hostname string, port uint32) []byte {
	pingHostString := "MC|PingHost"

	buff := make([]byte, 0)
	buff = append(buff, 0xFE, 0x01, 0xFA)

	buff = binary.BigEndian.AppendUint16(buff, uint16(len(pingHostString)))
	for _, val := range utf16.Encode([]rune(pingHostString)) {
		buff = binary.BigEndian.AppendUint16(buff, val)
	}

	hostnameBytes := utf16.Encode([]rune(hostname))
	buff = binary.BigEndian.AppendUint16(buff, uint16(7+(len(hostnameBytes)*2)))
	buff = append(buff, 0x4A)
	buff = binary.BigEndian.AppendUint16(buff, uint16(len(hostname)))
	for _, val := range hostnameBytes {
		buff = binary.BigEndian.AppendUint16(buff, val)
	}
	buff = binary.BigEndian.AppendUint32(buff, port)
	return buff
}
