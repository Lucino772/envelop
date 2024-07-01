package steamcm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/Lucino772/envelop/pkg/steam/steamlang"
)

type PacketMsg interface {
	IsProto() bool
	GetMsgType() steamlang.EMsg
	GetTargetJobId() uint64
	GetSourceJobId() uint64
	GetData() []byte
}

func ParseIncomingPacket(data []byte) (PacketMsg, error) {
	if len(data) < 4 {
		return nil, errors.New("not enough data")
	}
	t := steamlang.EMsg(binary.LittleEndian.Uint32(data[:4]))
	rd := bytes.NewBuffer(data)

	switch t {
	case steamlang.EMsg_ChannelEncryptRequest, steamlang.EMsg_ChannelEncryptResponse, steamlang.EMsg_ChannelEncryptResult:
		return parseBasicPacket(rd)
	}

	if isProtobuf(uint64(t)) {
		return parseProtobufPacket(rd)
	} else {
		return parseExtendedPacket(rd)
	}
}

func isProtobuf(m uint64) bool {
	return (m & uint64(PROTO_MASK)) > 0
}

func parseBasicPacket(r io.Reader) (*basicPacketClientMsg, error) {
	pkt := &basicPacketClientMsg{
		header: new(StdHeader),
	}
	if _, err := pkt.header.ReadFrom(r); err != nil {
		return nil, err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	pkt.data = data
	return pkt, nil
}

func parseExtendedPacket(r io.Reader) (*extendedPacketClientMsg, error) {
	pkt := &extendedPacketClientMsg{
		header: new(ExtHeader),
	}
	if _, err := pkt.header.ReadFrom(r); err != nil {
		return nil, err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	pkt.data = data
	return pkt, nil
}

func parseProtobufPacket(r io.Reader) (*protobufPacketClientMsg, error) {
	pkt := &protobufPacketClientMsg{
		header: new(ProtoHeader),
	}
	if _, err := pkt.header.ReadFrom(r); err != nil {
		return nil, err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	pkt.data = data
	return pkt, nil
}

type basicPacketClientMsg struct {
	header *StdHeader
	data   []byte
}

func (p *basicPacketClientMsg) IsProto() bool {
	return false
}

func (p *basicPacketClientMsg) GetMsgType() steamlang.EMsg {
	return p.header.MsgType
}

func (p *basicPacketClientMsg) GetTargetJobId() uint64 {
	return p.header.TargetJobId
}

func (p *basicPacketClientMsg) GetSourceJobId() uint64 {
	return p.header.SourceJobId
}

func (p *basicPacketClientMsg) GetData() []byte {
	return p.data
}

type extendedPacketClientMsg struct {
	header *ExtHeader
	data   []byte
}

func (p *extendedPacketClientMsg) IsProto() bool {
	return false
}

func (p *extendedPacketClientMsg) GetMsgType() steamlang.EMsg {
	return p.header.MsgType
}

func (p *extendedPacketClientMsg) GetTargetJobId() uint64 {
	return p.header.TargetJobId
}

func (p *extendedPacketClientMsg) GetSourceJobId() uint64 {
	return p.header.SourceJobId
}

func (p *extendedPacketClientMsg) GetData() []byte {
	return p.data
}

type protobufPacketClientMsg struct {
	header *ProtoHeader
	data   []byte
}

func (p *protobufPacketClientMsg) IsProto() bool {
	return true
}

func (p *protobufPacketClientMsg) GetMsgType() steamlang.EMsg {
	return p.header.MsgType
}

func (p *protobufPacketClientMsg) GetTargetJobId() uint64 {
	return 0
}

func (p *protobufPacketClientMsg) GetSourceJobId() uint64 {
	return 0
}

func (p *protobufPacketClientMsg) GetData() []byte {
	return p.data
}
