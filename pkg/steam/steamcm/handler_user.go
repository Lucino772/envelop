package steamcm

import (
	"log"
	"math"
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type steamUserHandler struct{}

func NewUserHandler() *steamUserHandler {
	return &steamUserHandler{}
}

func (handler *steamUserHandler) Register(handlers map[steamlang.EMsg]func(*Packet) ([]Event, error)) {
	handlers[steamlang.EMsg_ClientLogOnResponse] = handler.handleClientLogOnresponse
	handlers[steamlang.EMsg_ClientSessionToken] = handler.handleClientSessionToken
}

func (handler *steamUserHandler) LogInAnonymously(conn Connection) (*steampb.CMsgClientLogonResponse, error) {
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
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}

	return waitForJob[*steampb.CMsgClientLogonResponse](conn, math.MaxUint64, time.Second*30)
}

func (handler *steamUserHandler) handleClientLogOnresponse(packet *Packet) ([]Event, error) {
	if !packet.IsProto() {
		return nil, nil
	}

	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientLogonResponse]{
		Body: new(steampb.CMsgClientLogonResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.header.GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}

func (handler *steamUserHandler) handleClientSessionToken(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientSessionToken]{
		Body: new(steampb.CMsgClientSessionToken),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	log.Println("Session Token:", decoder.Body)
	return nil, nil
}
