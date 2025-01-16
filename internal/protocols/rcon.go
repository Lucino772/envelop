package protocols

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync/atomic"
)

var (
	ErrRconAuthFailed    = errors.New("rcon authentication failed")
	ErrRconInvalidPacket = errors.New("rcon invalid packet")
)

const (
	rconAuthRequest     int32 = 3
	rconExecCommand     int32 = 2
	rconAuthResponse    int32 = 2
	rconCommandResponse int32 = 0
)

func SendRcon(ctx context.Context, host string, port uint16, password string, command string) (string, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var client = rconClient{inner: conn}
	if err := client.authenticate(password); err != nil {
		return "", err
	}
	return client.exec(command)
}

type rconClient struct {
	inner     net.Conn
	requestId atomic.Int32
}

func (client *rconClient) authenticate(password string) error {
	var reqId = client.requestId.Add(1)
	if err := client.send(reqId, rconAuthRequest, []byte(password)); err != nil {
		return err
	}
	var buffer [4086]byte
	pktId, pktType, _, err := client.recv(buffer[:])
	if err != nil {
		return err
	}
	if pktType != rconAuthResponse {
		return ErrRconInvalidPacket
	}
	if pktId == -1 {
		return ErrRconAuthFailed
	}
	return nil
}

func (client *rconClient) exec(command string) (string, error) {
	var reqId = client.requestId.Add(1)
	if err := client.send(reqId, rconExecCommand, []byte(command)); err != nil {
		return "", err
	}

	// TODO: Implement packet fragmentation
	var buff [4086]byte
	pktId, pktType, nbytes, err := client.recv(buff[:])
	if err != nil {
		return "", err
	}
	if pktType != rconCommandResponse || pktId != reqId {
		return "", ErrRconInvalidPacket
	}
	return string(buff[:nbytes]), nil
}

func (client *rconClient) send(reqId int32, reqType int32, data []byte) error {
	var request = make([]byte, 14+len(data))
	binary.LittleEndian.PutUint32(request[0:4], uint32(10+len(data)))
	binary.LittleEndian.PutUint32(request[4:8], uint32(reqId))
	binary.LittleEndian.PutUint32(request[8:12], uint32(reqType))
	copy(request[12:], data)
	_, err := client.inner.Write(request)
	return err
}

func (client *rconClient) recv(buffer []byte) (int32, int32, int, error) {
	var header [12]byte
	if _, err := client.inner.Read(header[:]); err != nil {
		return -1, -1, -1, err
	}
	var pktSize = int32(binary.LittleEndian.Uint32(header[0:4]))
	var response = make([]byte, pktSize-8)
	if _, err := io.ReadFull(client.inner, response); err != nil {
		return -1, -1, -1, err
	}
	var pktId = int32(binary.LittleEndian.Uint32(header[4:8]))
	var pktType = int32(binary.LittleEndian.Uint32(header[8:12]))
	var nbytes = copy(buffer, response[:len(response)-2])
	return pktId, pktType, nbytes, nil
}
