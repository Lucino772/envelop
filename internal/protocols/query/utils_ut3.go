package query

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
)

var (
	ErrInvalidUT3Header = errors.New("invalid ut3 header")
)

func getUT3ChallengeToken(conn io.ReadWriter, sessionId int32) (int32, error) {
	var request = [7]byte{0xFE, 0xFD, 0x09}
	binary.BigEndian.PutUint32(request[3:], uint32(sessionId))
	if _, err := conn.Write(request[:]); err != nil {
		return -1, err
	}

	if err := readUT3Header(conn, sessionId, 9); err != nil {
		return -1, err
	}
	var buff bufio.Reader
	buff.Reset(conn)
	tokenString, err := buff.ReadString(0x00)
	if err != nil {
		return -1, err
	}
	token, err := strconv.ParseInt(tokenString[:len(tokenString)-1], 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(token), err
}

func sendUT3Query(conn io.ReadWriter, sessionId int32, token int32) error {
	var request = [15]byte{0xFE, 0xFD, 0x00}
	binary.BigEndian.PutUint32(request[3:], uint32(sessionId))
	binary.BigEndian.PutUint32(request[7:], uint32(token))
	if _, err := conn.Write(request[:]); err != nil {
		return err
	}
	return readUT3Header(conn, sessionId, 0)
}

func readUT3Header(rd io.Reader, sessionId int32, pktType byte) error {
	var header [5]byte
	if _, err := rd.Read(header[:]); err != nil {
		return err
	}
	var sessId = int32(binary.BigEndian.Uint32(header[1:]))
	if header[0] != pktType && sessId != sessionId {
		return ErrInvalidUT3Header
	}
	return nil
}
