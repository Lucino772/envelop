package steamcm

import (
	"log"
	"math"
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steammsg"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type SteamUserHandler struct{}

func NewUserHandler() *SteamUserHandler {
	return &SteamUserHandler{}
}

func (handler *SteamUserHandler) Register(handlers map[steamlang.EMsg]func(*steammsg.Packet) ([]Event, error)) {
	handlers[steamlang.EMsg_ClientLogOnResponse] = handler.handleClientLogOnresponse
	handlers[steamlang.EMsg_ClientSessionToken] = handler.handleClientSessionToken
}

func (handler *SteamUserHandler) LogInAnonymously(conn Connection) (*steampb.CMsgClientLogonResponse, error) {
	audId := steam.NewInstanceSteamId(0, steam.Instance_All, steamlang.EUniverse_Public, steamlang.EAccountType_AnonUser)
	header := steammsg.NewProtoHeader(steamlang.EMsg_ClientLogon)
	header.Proto.ClientSessionid = proto.Int32(0)
	header.Proto.Steamid = proto.Uint64(uint64(audId))
	body := &steampb.CMsgClientLogon{
		ProtocolVersion: proto.Uint32(65580),
		ClientOsType:    proto.Uint32(20),
		ClientLanguage:  proto.String("english"),
		CellId:          proto.Uint32(0),
	}
	packet, err := steammsg.EncodePacket(header, body, nil)
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}

	return waitForJob[*steampb.CMsgClientLogonResponse](conn, math.MaxUint64, time.Second*30)
}

func (handler *SteamUserHandler) handleClientLogOnresponse(packet *steammsg.Packet) ([]Event, error) {
	if !packet.IsProto() {
		return nil, nil
	}

	var decoder = &steammsg.ProtoPacketDecoder[*steampb.CMsgClientLogonResponse]{
		Body: new(steampb.CMsgClientLogonResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.Header().GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}

func (handler *SteamUserHandler) handleClientSessionToken(packet *steammsg.Packet) ([]Event, error) {
	var decoder = &steammsg.ProtoPacketDecoder[*steampb.CMsgClientSessionToken]{
		Body: new(steampb.CMsgClientSessionToken),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	log.Println("Session Token:", decoder.Body)
	return nil, nil
}
