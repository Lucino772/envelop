package steamcm

import "github.com/Lucino772/envelop/pkg/steam"

type EventType = uint

const (
	EventType_Incoming EventType = 0
	EventType_Outgoing EventType = 1
	EventType_State    EventType = 2
)

type Event struct {
	Type    EventType
	Payload any
}

func MakeEvent(etype EventType, payload any) Event {
	return Event{Type: etype, Payload: payload}
}

func (event *Event) WithPayload(payload any) Event {
	return Event{Type: event.Type, Payload: payload}
}

type EventChannelEncrypted struct{}

type EventDataReceived struct {
	Data []byte
}
type EventDataToSend struct {
	Data []byte
}
type EventPacketReceived struct {
	Packet *Packet
}
type EventPacketTosend struct {
	Packet *Packet
}

type EventCallback struct {
	JobId   steam.JobId
	Payload any
}

type Layer interface {
	ProcessIncoming([]Event) ([]Event, error)
	ProcessOutgoing([]Event) ([]Event, error)
}
