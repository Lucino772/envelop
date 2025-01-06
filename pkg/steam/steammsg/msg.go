package steammsg

import "github.com/Lucino772/envelop/pkg/steam/steamlang"

const (
	PROTO_MASK uint32 = 0x80000000
	EMSG_MASK  uint32 = ^PROTO_MASK
)

func GetMsg(msg uint32) steamlang.EMsg {
	return steamlang.EMsg(msg & EMSG_MASK)
}

func IsProtobuf(m uint32) bool {
	return (m & PROTO_MASK) > 0
}

func MakeMsg(msg steamlang.EMsg, protobuf bool) steamlang.EMsg {
	if protobuf {
		return steamlang.EMsg(uint32(msg) | PROTO_MASK)
	}
	return msg
}
