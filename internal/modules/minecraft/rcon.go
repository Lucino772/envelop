package minecraft

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
)

const (
	PACKET_LOGIN    = 3
	PACKET_COMMAND  = 2
	PACKET_RESPONSE = 0
)

func RconSend(ctx context.Context, host string, port int, password string, command string) (string, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = makeRequest(conn, PACKET_LOGIN, PACKET_COMMAND, password)
	if err != nil {
		return "", err
	}
	return makeRequest(conn, PACKET_COMMAND, PACKET_RESPONSE, command)
}

func makeRequest(wr io.ReadWriter, wType int32, rType int32, paylaod string) (string, error) {
	wPktId, err := writeRconPacket(wr, wType, paylaod)
	if err != nil {
		return "", err
	}
	rPktId, rPktType, response, err := readRconPacket(wr)
	if err != nil {
		return "", err
	}
	if rPktId != wPktId || rType != rPktType {
		return "", errors.New("invalid packet")
	}
	return string(response[0 : len(response)-2]), nil
}

func writeRconPacket(w io.Writer, typ int32, payload string) (int32, error) {
	var pktId int32 = rand.Int31()
	toWrite := []any{
		int32(10 + len(payload)),
		int32(pktId),
		int32(typ),
	}

	var buf bytes.Buffer
	for _, d := range toWrite {
		if err := binary.Write(&buf, binary.LittleEndian, d); err != nil {
			return -1, err
		}
	}
	payloadBytes := []byte(payload)
	payloadBytes = append(payloadBytes, 0, 0)
	if _, err := buf.Write(payloadBytes); err != nil {
		return -1, err
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		return -1, err
	}
	return pktId, nil
}

func readRconPacket(r io.Reader) (pktId int32, pktType int32, data []byte, err error) {
	rd := bufio.NewReader(r)

	var pktSize int32
	if err = binary.Read(rd, binary.LittleEndian, &pktSize); err != nil {
		return 0, 0, nil, err
	}
	if err = binary.Read(rd, binary.LittleEndian, &pktId); err != nil {
		return 0, 0, nil, err
	}
	if err = binary.Read(rd, binary.LittleEndian, &pktType); err != nil {
		return 0, 0, nil, err
	}

	data = make([]byte, pktSize-8)
	nr, err := rd.Read(data)
	if err != nil {
		return 0, 0, nil, err
	}
	if nr != len(data) {
		return 0, 0, nil, io.ErrUnexpectedEOF
	}

	return pktId, pktType, data, nil
}
