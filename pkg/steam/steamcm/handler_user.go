package steamcm

import (
	"log"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type UserHandler struct {
	conn PacketConnection
}

func (handler *UserHandler) Register(handlers map[steamlang.EMsg]func(*Packet) error) {
	handlers[steamlang.EMsg_ClientLogOnResponse] = handler.handleClientLogOnResponse
}

func (handler *UserHandler) LogInAnonymously() error {
	audId := steam.NewInstanceSteamId(0, steam.Instance_All, steamlang.EUniverse_Public, steamlang.EAccountType_AnonUser)
	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ClientLogon)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.ClientSessionid = proto.Int32(0)
	header.Proto.Steamid = proto.Uint64(uint64(audId))
	encoder.Body = &steampb.CMsgClientLogon{
		ProtocolVersion: proto.Uint32(65580),
		ClientOsType:    proto.Uint32(20),
		ClientLanguage:  proto.String("english"),
		CellId:          proto.Uint32(0),
	}
	packet, err := encoder.Encode()
	if err != nil {
		return err
	}
	return handler.conn.SendPacket(packet)
}

func (handler *UserHandler) handleClientLogOnResponse(packet *Packet) error {
	if packet.IsProto() {
		var decoder = &ProtoPacketDecoder[*steampb.CMsgClientLogonResponse]{
			Body: new(steampb.CMsgClientLogonResponse),
		}
		if err := decoder.Decode(packet); err != nil {
			return err
		}
		log.Println("Login Result (Proto):", decoder.Body.GetEresult())
	} else {
		var decoder = &PacketDecoder[*MsgClientLogOnResponse]{
			Body: new(MsgClientLogOnResponse),
		}
		if err := decoder.Decode(packet); err != nil {
			return err
		}
		log.Println("Login Result:", decoder.Body.Result)
	}
	return nil
}
