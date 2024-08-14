package rcon

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

var (
	PACKET_LOGIN    = []byte{0x03, 0x00, 0x00, 0x00}
	PACKET_COMMAND  = []byte{0x02, 0x00, 0x00, 0x00}
	PACKET_RESPONSE = []byte{0x00, 0x00, 0x00, 0x00}
)

func Send(ctx context.Context, host string, port uint16, password string, command string) (string, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = sendPacket(conn, PACKET_LOGIN, []byte(password), PACKET_COMMAND)
	if err != nil {
		return "", err
	}
	response, err := sendPacket(conn, PACKET_COMMAND, []byte(command), PACKET_RESPONSE)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func sendPacket(conn net.Conn, reqType []byte, data []byte, respType []byte) ([]byte, error) {
	request := make([]byte, 14+len(data))
	binary.LittleEndian.PutUint32(request[0:4], uint32(10+len(data)))
	if _, err := rand.Read(request[4:8]); err != nil {
		return nil, err
	}
	copy(request[8:12], reqType)
	copy(request[12:], []byte(data))
	if _, err := conn.Write(request); err != nil {
		return nil, err
	}

	var header [12]byte
	if _, err := conn.Read(header[:]); err != nil {
		return nil, err
	}
	pktSize := int32(binary.LittleEndian.Uint32(header[0:4]))

	response := make([]byte, pktSize-8)
	if _, err := io.ReadFull(conn, response); err != nil {
		return nil, err
	}

	if !bytes.Equal(request[4:8], header[4:8]) || !bytes.Equal(header[8:12], respType) {
		return nil, errors.New("invalid packet")
	}

	return response[:len(response)-2], nil
}
