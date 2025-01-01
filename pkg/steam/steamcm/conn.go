package steamcm

import (
	"bytes"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
)

type Connection interface {
	SendPacket(*Packet) error
	GetNextJobId() steam.JobId
	RegisterJob(steam.JobId, func(any))
}

type SteamConnection struct {
	layer      Layer
	mu         sync.Mutex
	cond       *sync.Cond
	dataToSend bytes.Buffer
	startTime  time.Time

	jobIdSequence atomic.Uint32
	jobs          map[steam.JobId]func(any)
}

func NewSteamConnection(layer Layer) *SteamConnection {
	conn := &SteamConnection{
		layer:     layer,
		startTime: time.Now(),
		jobs:      make(map[steam.JobId]func(any)),
	}
	conn.cond = sync.NewCond(&conn.mu)
	return conn
}

func (conn *SteamConnection) GetNextJobId() steam.JobId {
	conn.jobIdSequence.Add(1)
	var jobId steam.JobId = 0
	jobId.SetBoxId(0)
	jobId.SetProcessId(0)
	jobId.SetSequentialCount(conn.jobIdSequence.Load())
	jobId.SetStartTime(conn.startTime)
	return jobId
}

func (conn *SteamConnection) RegisterJob(jobId steam.JobId, callback func(any)) {
	conn.jobs[jobId] = callback
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

func (conn *SteamConnection) handleEvent(event Event) error {
	switch ev := event.Payload.(type) {
	case EventPacketReceived:
		log.Println("Following packet was not handled:", ev.Packet.MsgType())
	case EventDataToSend:
		if _, err := conn.dataToSend.Write(ev.Data); err != nil {
			return err
		}
		conn.cond.Signal()
	case EventChannelEncrypted:
		log.Println("Channel successfully encrypted !")
	case EventCallback:
		if callback, ok := conn.jobs[ev.JobId]; ok {
			callback(ev.Payload)
		}
	}
	return nil
}
