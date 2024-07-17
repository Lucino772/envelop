package steamcm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"

	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"google.golang.org/protobuf/proto"
)

type Packet struct {
	Header PacketHeader
	buf    *bytes.Buffer
}

func NewPacket(emsg steamlang.EMsg) *Packet {
	return &Packet{
		Header: &StdHeader{
			MsgType:     emsg,
			TargetJobId: math.MaxUint64,
			SourceJobId: math.MaxUint64,
		},
	}
}

func NewExtPacket(emsg steamlang.EMsg) *Packet {
	return &Packet{
		Header: &ExtHeader{
			MsgType:     emsg,
			TargetJobId: math.MaxUint64,
			SourceJobId: math.MaxUint64,
		},
	}
}

func NewProtoPacket(emsg steamlang.EMsg) *Packet {
	return &Packet{
		Header: &ProtoHeader{
			MsgType: emsg,
		},
	}
}

func ParsePacket(data []byte) (*Packet, error) {
	if len(data) < 4 {
		return nil, errors.New("not enough data")
	}
	emsg := steamlang.EMsg(binary.LittleEndian.Uint32(data[:4]))

	var pkt Packet
	switch emsg {
	case steamlang.EMsg_ChannelEncryptRequest,
		steamlang.EMsg_ChannelEncryptResponse,
		steamlang.EMsg_ChannelEncryptResult:
		pkt.Header = new(StdHeader)
	default:
		if IsProtobuf(uint32(emsg)) {
			pkt.Header = new(ProtoHeader)
		} else {
			pkt.Header = new(ExtHeader)
		}
	}

	buf := bytes.NewBuffer(data)
	if _, err := pkt.Header.ReadFrom(buf); err != nil {
		return nil, err
	}
	pkt.buf = buf
	return &pkt, nil
}

func (p *Packet) MsgType() steamlang.EMsg {
	return p.Header.GetMsgType()
}

func (p *Packet) UnmarshalBody(body io.ReaderFrom) ([]byte, error) {
	_, err := body.ReadFrom(p.buf)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(p.buf)
}

func (p *Packet) UnmarshalProtoBody(body proto.Message) error {
	data, err := io.ReadAll(p.buf)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, body)
}

func (p *Packet) MarshalBody(body io.WriterTo) error {
	p.buf = bytes.NewBuffer([]byte{})
	_, err := body.WriteTo(p.buf)
	return err
}

func (p *Packet) MarshalProtoBody(body proto.Message) error {
	data, err := proto.Marshal(body)
	if err != nil {
		return err
	}
	p.buf = bytes.NewBuffer(data)
	return err
}

func (p *Packet) Write(data []byte) (int, error) {
	return p.buf.Write(data)
}

func (p *Packet) Bytes() []byte {
	var buf bytes.Buffer
	if _, err := p.Header.WriteTo(&buf); err != nil {
		return nil
	}
	if _, err := p.buf.WriteTo(&buf); err != nil {
		return nil
	}
	return buf.Bytes()
}
