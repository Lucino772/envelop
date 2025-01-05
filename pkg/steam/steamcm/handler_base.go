package steamcm

import (
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
)

type SteamBaseHandler struct {
}

func NewSteamBaseHandler() *SteamBaseHandler {
	return &SteamBaseHandler{}
}

func (handler *SteamBaseHandler) Register(handlers map[steamlang.EMsg]func(*Packet) ([]Event, error)) {
	handlers[steamlang.EMsg_ClientServerUnavailable] = handler.handleServerUnavailable
	handlers[steamlang.EMsg_ClientCMList] = handler.handleCMList
	handlers[steamlang.EMsg_ClientSessionToken] = handler.handleSessionToken

}

func (handler *SteamBaseHandler) handleServerUnavailable(packet *Packet) ([]Event, error) {
	var decoder = &PacketDecoder[*MsgClientServerUnavailable]{
		Body: new(MsgClientServerUnavailable),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	// TODO: Close connection
	return nil, nil
}

func (handler *SteamBaseHandler) handleCMList(_ *Packet) ([]Event, error) {
	// TODO: Read and update CM list
	return nil, nil
}

func (handler *SteamBaseHandler) handleSessionToken(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientSessionToken]{
		Body: new(steampb.CMsgClientSessionToken),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}

	// TODO: Set internal session token
	// body.GetToken()
	return nil, nil
}
