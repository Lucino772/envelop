package steamcm

import (
	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
)

type sessionLayer struct {
	steamId   *steam.SteamId
	sessionId *int32
}

func NewSessionLayer() *sessionLayer {
	return &sessionLayer{
		steamId:   nil,
		sessionId: nil,
	}
}

func (layer *sessionLayer) ProcessIncoming(events []Event) ([]Event, error) {
	processedEvents := make([]Event, 0)
	for _, event := range events {
		if event.Type != EventType_Incoming {
			processedEvents = append(processedEvents, event)
			continue
		}

		switch payload := event.Payload.(type) {
		case EventPacketReceived:
			_events, err := layer.handleIncomingPacket(payload.Packet)
			if err != nil {
				return nil, err
			}
			processedEvents = append(processedEvents, _events...)
		default:
			processedEvents = append(processedEvents, event)
		}
	}
	return processedEvents, nil
}

func (layer *sessionLayer) ProcessOutgoing(events []Event) ([]Event, error) {
	processedEvents := make([]Event, 0)
	for _, event := range events {
		if event.Type != EventType_Outgoing {
			processedEvents = append(processedEvents, event)
			continue
		}

		switch payload := event.Payload.(type) {
		case EventPacketTosend:
			packet := payload.Packet
			if layer.steamId != nil {
				packet.Header().SetSteamId(layer.steamId)
			}
			if layer.sessionId != nil {
				packet.Header().SetSessionId(layer.sessionId)
			}
			processedEvents = append(
				processedEvents,
				event.WithPayload(EventPacketTosend{Packet: packet}),
			)
		default:
			processedEvents = append(processedEvents, event)
		}
	}
	return processedEvents, nil
}

func (layer *sessionLayer) handleIncomingPacket(packet *Packet) ([]Event, error) {
	events := make([]Event, 0)

	switch packet.MsgType() {
	case steamlang.EMsg_ClientLogOnResponse:
		if !packet.IsProto() {
			return nil, nil
		}
		var decoder = &ProtoPacketDecoder[*steampb.CMsgClientLogonResponse]{
			Body: new(steampb.CMsgClientLogonResponse),
		}
		if err := decoder.Decode(packet); err != nil {
			return nil, err
		}
		if decoder.Body.GetEresult() == int32(steamlang.EResult_OK) {
			layer.steamId = packet.header.GetSteamId()
			layer.sessionId = packet.Header().GetSessionId()
		} else {
			layer.steamId = nil
			layer.sessionId = nil
		}
	case steamlang.EMsg_ClientLoggedOff:
		layer.steamId = nil
		layer.sessionId = nil
		if !packet.IsProto() {
			return nil, nil
		}
		var decoder = &ProtoPacketDecoder[*steampb.CMsgClientLoggedOff]{
			Body: new(steampb.CMsgClientLoggedOff),
		}
		if err := decoder.Decode(packet); err != nil {
			return nil, err
		}
	}

	events = append(
		events,
		MakeEvent(EventType_Incoming, EventPacketReceived{Packet: packet}),
	)
	return events, nil
}
