package steamcm

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync/atomic"
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"golang.org/x/sync/errgroup"
)

type Connection interface {
	SendPacket(*Packet) error
	GetNextJobId() steam.JobId
	RegisterJob(steam.JobId, func(any))
}

type SteamConnection struct {
	layer         Layer
	dataToSend    io.ReadWriter
	startTime     time.Time
	jobIdSequence atomic.Uint32
	jobs          map[steam.JobId]func(any)
	queuedEvents  chan Event
	errg          *errgroup.Group
	readyChan     chan struct{}
}

func NewSteamConnection(handlers ...Handler) *SteamConnection {
	conn := &SteamConnection{
		layer: MakeLayerStack(
			NewTCPLayer(),
			NewEncryptedLayer(steamlang.EUniverse_Public),
			NewPacketLayer(),
			NewSessionLayer(),
			MakeDispatchLayer(handlers...),
		),
		startTime:    time.Now(),
		jobs:         make(map[steam.JobId]func(any)),
		queuedEvents: make(chan Event),
		errg:         new(errgroup.Group),
		dataToSend:   NewBuffer(),
		readyChan:    make(chan struct{}),
	}
	return conn
}

func (conn *SteamConnection) Connect() error {
	s := new(Servers)
	if err := s.Update(); err != nil {
		return err
	}
	server := s.Records()[0]

	tcpConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		return err
	}
	return conn.netLoop(tcpConn)
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

func (conn *SteamConnection) SendPacket(packet *Packet) error {
	conn.queuedEvents <- MakeEvent(EventType_Outgoing, EventPacketTosend{Packet: packet})
	return nil
}

func (conn *SteamConnection) WaitReady(timeout time.Duration) error {
	select {
	case <-conn.readyChan:
		return nil
	case <-time.After(timeout):
		return context.DeadlineExceeded
	}
}

func (conn *SteamConnection) netLoop(transportConn net.Conn) error {
	conn.errg.Go(func() error {
		defer log.Println("client: done reading")
		for {
			buff := make([]byte, 1024)
			log.Println("client: reading...")
			nr, err := transportConn.Read(buff)
			if err != nil {
				return err
			}
			log.Println("client: recv data", nr)
			conn.queuedEvents <- MakeEvent(
				EventType_Incoming,
				EventDataReceived{Data: buff[:nr]},
			)
		}
	})
	conn.errg.Go(func() error {
		defer log.Println("client: done writing")
		for {
			buff := make([]byte, 1024)
			log.Println("client: waiting for data to send")
			nr, err := conn.dataToSend.Read(buff)
			if err != nil {
				return err
			}
			log.Println("client: send data", len(buff[:nr]))
			if _, err := transportConn.Write(buff[:nr]); err != nil {
				return err
			}
		}
	})
	conn.errg.Go(func() error {
		defer log.Println("client: done processing")
		for event := range conn.queuedEvents {
			err := conn.processEvent(event)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (conn *SteamConnection) processEvent(event Event) error {
	events := make([]Event, 0)
	switch event.Type {
	case EventType_Incoming:
		_events, err := conn.layer.ProcessIncoming([]Event{event})
		if err != nil {
			return err
		}
		events = append(events, _events...)
	case EventType_Outgoing:
		_events, err := conn.layer.ProcessOutgoing([]Event{event})
		if err != nil {
			return err
		}
		events = append(events, _events...)
	default:
		events = append(events, event)
	}
	return conn.handleEvents(events)
}

func (conn *SteamConnection) handleEvents(events []Event) error {
	for _, event := range events {
		switch payload := event.Payload.(type) {
		case EventDataToSend:
			if _, err := conn.dataToSend.Write(payload.Data); err != nil {
				return err
			}
		case EventPacketReceived:
			log.Println("Following packet was not handled:", payload.Packet.MsgType())
		case EventChannelEncrypted:
			close(conn.readyChan)
		case EventCallback:
			if callback, ok := conn.jobs[payload.JobId]; ok {
				callback(payload.Payload)
			}
		default:
			log.Println("Unhandled event", event)
		}
	}
	return nil
}
