package steamcm

import (
	"bytes"
	"log"
	"sync"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
)

type Connection interface {
	HandlePacket(*Packet) error
	SendPacket(*Packet) error
	AddHandler(steamlang.EMsg, func(*Packet) error)
}

type SteamConnection struct {
	layer      Layer
	handlers   map[steamlang.EMsg]func(*Packet) error
	mu         sync.Mutex
	cond       *sync.Cond
	dataToSend bytes.Buffer

	steamId   *steam.SteamId
	sessionId *int32
}

func NewSteamConnection(layer Layer) *SteamConnection {
	conn := &SteamConnection{
		layer:    layer,
		handlers: make(map[steamlang.EMsg]func(*Packet) error),
	}
	conn.cond = sync.NewCond(&conn.mu)
	return conn
}

func (conn *SteamConnection) SetSteamId(id *steam.SteamId) {
	conn.steamId = id
}

func (conn *SteamConnection) SetSessionId(id *int32) {
	conn.sessionId = id
}

func (conn *SteamConnection) ProcessBytes(data []byte) error {
	events, err := conn.layer.ProcessIncoming([]Event{
		MakeEvent(EventType_Incoming, EventDataReceived{Data: data}),
	})
	if err != nil {
		return err
	}
	for _, event := range events {
		if err := conn.handleEvent(event); err != nil {
			return err
		}
	}
	return nil
}

func (conn *SteamConnection) Read(data []byte) (int, error) {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.dataToSend.Len() == 0 {
		conn.cond.Wait()
	}
	return conn.dataToSend.Read(data)
}

func (conn *SteamConnection) AddHandler(msg steamlang.EMsg, handler func(*Packet) error) {
	conn.handlers[msg] = handler
}

func (conn *SteamConnection) SendPacket(packet *Packet) error {
	events, err := conn.layer.ProcessOutgoing([]Event{
		MakeEvent(EventType_Outgoing, EventPacketTosend{Packet: packet}),
	})
	if err != nil {
		return err
	}
	for _, event := range events {
		if err := conn.handleEvent(event); err != nil {
			return err
		}
	}
	return nil
}

func (conn *SteamConnection) HandlePacket(packet *Packet) error {
	if handler, ok := conn.handlers[packet.MsgType()]; ok {
		return handler(packet)
	}
	return nil
}

func (conn *SteamConnection) handleEvent(event Event) error {
	switch ev := event.Payload.(type) {
	case EventPacketReceived:
		return conn.HandlePacket(ev.Packet)
	case EventDataToSend:
		if _, err := conn.dataToSend.Write(ev.Data); err != nil {
			return err
		}
		conn.cond.Signal()
	case EventChannelEncrypted:
		log.Println("Channel successfully encrypted !")
	case EventLogOnSuccess:
		log.Println("Successfully logged-on !")
	case EventLogOnError:
		log.Println("Failed to log-in !")
	}
	return nil
}
