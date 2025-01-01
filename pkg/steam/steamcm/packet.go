package steamcm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"

	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type Packet struct {
	header PacketHeader
	buf    *bytes.Buffer
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
		pkt.header = new(StdHeader)
	default:
		if IsProtobuf(uint32(emsg)) {
			pkt.header = new(ProtoHeader)
		} else {
			pkt.header = new(ExtHeader)
		}
	}

	buf := bytes.NewBuffer(data)
	if _, err := pkt.header.ReadFrom(buf); err != nil {
		return nil, err
	}
	pkt.buf = buf
	return &pkt, nil
}

func (p *Packet) MsgType() steamlang.EMsg {
	return p.header.GetMsgType()
}

func (p *Packet) Header() PacketHeader {
	return p.header
}

func (p *Packet) IsProto() bool {
	if _, ok := p.header.(*ProtoHeader); ok {
		return true
	}
	return false
}

func (p *Packet) Bytes() []byte {
	var buf bytes.Buffer
	if _, err := p.header.WriteTo(&buf); err != nil {
		return nil
	}
	if _, err := p.buf.WriteTo(&buf); err != nil {
		return nil
	}
	return buf.Bytes()
}

type PacketDecoder[Body_T io.ReaderFrom] struct {
	Body    Body_T
	Payload []byte
}

func (decoder *PacketDecoder[T]) Decode(packet *Packet) error {
	tempBuf := bytes.NewBuffer(packet.buf.Bytes())

	_, err := decoder.Body.ReadFrom(tempBuf)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(tempBuf)
	if err != nil {
		return err
	}
	decoder.Payload = data
	return nil
}

type ProtoPacketDecoder[Body_T proto.Message] struct {
	Body Body_T
}

func (decoder *ProtoPacketDecoder[T]) Decode(packet *Packet) error {
	return proto.Unmarshal(packet.buf.Bytes(), decoder.Body)
}

type PacketEncoder struct {
	Header PacketHeader
	Body   any
	Data   *bytes.Buffer
}

func NewPacketEncoder(emsg steamlang.EMsg) *PacketEncoder {
	return &PacketEncoder{
		Header: &StdHeader{
			MsgType:     emsg,
			TargetJobId: math.MaxUint64,
			SourceJobId: math.MaxUint64,
		},
		Body: nil,
		Data: bytes.NewBuffer([]byte{}),
	}
}

func NewExtPacketEncoder(emsg steamlang.EMsg) *PacketEncoder {
	return &PacketEncoder{
		Header: &ExtHeader{
			MsgType:     emsg,
			TargetJobId: math.MaxUint64,
			SourceJobId: math.MaxUint64,
		},
		Body: nil,
		Data: bytes.NewBuffer([]byte{}),
	}
}

func NewProtoPacketEncoder(emsg steamlang.EMsg) *PacketEncoder {
	return &PacketEncoder{
		Header: &ProtoHeader{
			MsgType: emsg,
			Proto:   new(steampb.CMsgProtoBufHeader),
		},
		Body: nil,
		Data: nil,
	}
}

func (encoder *PacketEncoder) Encode() (*Packet, error) {
	if encoder.Body == nil {
		return nil, errors.New("missing packet body")
	}

	var pkt Packet
	pkt.header = encoder.Header
	if body, ok := encoder.Body.(proto.Message); ok {
		data, err := proto.Marshal(body)
		if err != nil {
			return nil, err
		}
		pkt.buf = bytes.NewBuffer(data)
	} else if body, ok := encoder.Body.(io.WriterTo); ok {
		pkt.buf = bytes.NewBuffer([]byte{})
		if _, err := body.WriteTo(pkt.buf); err != nil {
			return nil, err
		}
		if encoder.Data != nil {
			if _, err := encoder.Data.WriteTo(pkt.buf); err != nil {
				return nil, err
			}
		}
	} else {
		return nil, errors.New("incomptabile packet body")
	}
	return &pkt, nil
}
