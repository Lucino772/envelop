package steamcm

import (
	"log"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type steamUserHandler struct {
	conn Connection
}

func NewUserHandler(conn Connection) *steamUserHandler {
	handler := &steamUserHandler{conn}
	handler.conn.AddHandler(steamlang.EMsg_ClientLogOnResponse, handler.handleClientLogOnResponse)
	handler.conn.AddHandler(steamlang.EMsg_ClientLoggedOff, handler.handleClientLoggedOff)
	handler.conn.AddHandler(steamlang.EMsg_ClientSessionToken, handler.handleClientSessionToken)
	return handler
}

func (handler *steamUserHandler) LogInAnonymously() error {
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

func (handler *steamUserHandler) handleClientLogOnResponse(packet *Packet) error {
	if packet.IsProto() {
		var decoder = &ProtoPacketDecoder[*steampb.CMsgClientLogonResponse]{
			Body: new(steampb.CMsgClientLogonResponse),
		}
		if err := decoder.Decode(packet); err != nil {
			return err
		}
		log.Println("Login Result (Proto):", decoder.Body)
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

func (handler *steamUserHandler) handleClientLoggedOff(packet *Packet) error {
	if packet.IsProto() {
		var decoder = &ProtoPacketDecoder[*steampb.CMsgClientLoggedOff]{
			Body: new(steampb.CMsgClientLoggedOff),
		}
		if err := decoder.Decode(packet); err != nil {
			return err
		}
		log.Println("Logged Off Result (Proto):", decoder.Body.GetEresult())
	} else {
		var decoder = &PacketDecoder[*MsgClientLoggedOff]{
			Body: new(MsgClientLoggedOff),
		}
		if err := decoder.Decode(packet); err != nil {
			return err
		}
		log.Println("Logged Off Result:", decoder.Body.Result)
	}
	return nil
}

func (handler *steamUserHandler) handleClientSessionToken(packet *Packet) error {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientSessionToken]{
		Body: new(steampb.CMsgClientSessionToken),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}
	log.Println("Session Token:", decoder.Body)
	return nil
}
