package steamcm

import (
	"encoding/binary"
	"io"

	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type StdHeader struct {
	MsgType     steamlang.EMsg
	TargetJobId uint64
	SourceJobId uint64
}

func (h *StdHeader) ReadFrom(r io.Reader) (int64, error) {
	if err := binary.Read(r, binary.LittleEndian, h); err != nil {
		return -1, err
	}
	return int64(binary.Size(h)), nil
}

func (h *StdHeader) WriteTo(w io.Writer) (int64, error) {
	if err := binary.Write(w, binary.LittleEndian, h); err != nil {
		return -1, err
	}
	return int64(binary.Size(h)), nil
}

type ExtHeader struct {
	HeaderSize    byte
	HeaderVersion uint16
	MsgType       steamlang.EMsg
	TargetJobId   uint64
	SourceJobId   uint64
	HeaderCanary  byte
	SteamId       uint64
	SessionId     int32
}

func (h *ExtHeader) ReadFrom(r io.Reader) (int64, error) {
	if err := binary.Read(r, binary.LittleEndian, h); err != nil {
		return -1, err
	}
	return int64(binary.Size(h)), nil
}

func (h *ExtHeader) WriteTo(w io.Writer) (int64, error) {
	if err := binary.Write(w, binary.LittleEndian, h); err != nil {
		return -1, err
	}
	return int64(binary.Size(h)), nil
}

type ProtoHeader struct {
	MsgType   steamlang.EMsg
	HeaderLen int32
	Proto     steampb.CMsgProtoBufHeader
}

func (h *ProtoHeader) ReadFrom(r io.Reader) (int64, error) {
	var nbytes int64 = 0
	if err := binary.Read(r, binary.LittleEndian, &h.MsgType); err != nil {
		return nbytes, err
	}
	nbytes += int64(binary.Size(h.MsgType))
	if err := binary.Read(r, binary.LittleEndian, &h.HeaderLen); err != nil {
		return nbytes, err
	}
	nbytes += int64(binary.Size(h.HeaderLen))
	headerBytes := make([]byte, h.HeaderLen)
	if _, err := io.ReadFull(r, headerBytes); err != nil {
		return nbytes, err
	}
	nbytes += int64(len(headerBytes))
	if err := proto.Unmarshal(headerBytes, &h.Proto); err != nil {
		return nbytes, err
	}
	return nbytes, nil
}

func (h *ProtoHeader) WriteTo(w io.Writer) (int64, error) {
	var nbytes int64 = 0
	if err := binary.Write(w, binary.LittleEndian, h.MsgType); err != nil {
		return nbytes, err
	}
	nbytes += int64(binary.Size(h.MsgType))

	headerBytes, err := proto.Marshal(&h.Proto)
	if err != nil {
		return nbytes, err
	}
	headerLen := int32(len(headerBytes))

	if err := binary.Write(w, binary.LittleEndian, headerLen); err != nil {
		return nbytes, err
	}
	nbytes += int64(binary.Size(headerLen))

	if _, err := w.Write(headerBytes); err != nil {
		return nbytes, err
	}
	nbytes += int64(headerLen)
	return nbytes, nil
}
