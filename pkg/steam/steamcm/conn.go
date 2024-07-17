package steamcm

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

const (
	tcpConnectionMagic uint32 = 0x31305456
)

type Connection interface {
	SendPacket(packet []byte) error
	ReadPacket() ([]byte, error)
	Close() error
}

type TCPConnection struct {
	inner net.Conn
}

func (conn *TCPConnection) SendPacket(packet []byte) error {
	var pktHeader = struct {
		PktLen   uint32
		PktMagic uint32
	}{
		uint32(len(packet)),
		tcpConnectionMagic,
	}
	if err := binary.Write(conn.inner, binary.LittleEndian, pktHeader); err != nil {
		return err
	}
	_, err := conn.inner.Write(packet)
	return err
}

func (conn *TCPConnection) ReadPacket() ([]byte, error) {
	var pktHeader struct {
		PktLen   uint32
		PktMagic uint32
	}
	if err := binary.Read(conn.inner, binary.LittleEndian, &pktHeader); err != nil {
		return nil, err
	}
	if pktHeader.PktMagic != tcpConnectionMagic {
		return nil, errors.New("unknown packet")
	}
	pktData := make([]byte, pktHeader.PktLen)
	if _, err := io.ReadFull(conn.inner, pktData); err != nil {
		return nil, err
	}
	return pktData, nil
}

func (conn *TCPConnection) Close() error {
	return conn.inner.Close()
}
