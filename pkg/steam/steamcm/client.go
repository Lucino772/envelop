package steamcm

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"golang.org/x/sync/errgroup"
)

type CMClient struct {
	UserHandler *UserHandler

	inner          PacketConnection
	errg           *errgroup.Group
	packetHandlers map[steamlang.EMsg]func(*Packet) error
}

func NewCMClient() (*CMClient, error) {
	s := new(Servers)
	if err := s.Update(); err != nil {
		return nil, err
	}
	server := s.Records()[0]

	conn, err := NewTCPConnection(server.Host, server.Port)
	if err != nil {
		return nil, err
	}
	cm := &CMClient{
		inner:          NewEncryptedConnection(steamlang.EUniverse_Public, conn),
		errg:           &errgroup.Group{},
		packetHandlers: make(map[steamlang.EMsg]func(*Packet) error),
	}
	cm.UserHandler = &UserHandler{conn: cm}
	cm.UserHandler.Register(cm.packetHandlers)

	cm.errg.Go(cm.netLoop)
	return cm, nil
}

func (conn *CMClient) SendPacket(packet *Packet) error {
	return conn.inner.SendPacket(packet)
}

func (conn *CMClient) RecvPacket() (*Packet, error) {
	return conn.inner.RecvPacket()
}

func (conn *CMClient) Close() error {
	err := conn.inner.Close()
	if err := conn.errg.Wait(); err != nil {
		return err
	}
	return err
}

func (conn *CMClient) HandlePacket(packet *Packet) (*Packet, error) {
	log.Println("Handling packet:", packet.MsgType())

	packet, err := conn.inner.HandlePacket(packet)
	if err != nil || packet == nil {
		return nil, err
	}
	switch packet.MsgType() {
	case steamlang.EMsg_Multi:
		return nil, conn.handleMulti(packet)
	case steamlang.EMsg_ClientServerUnavailable:
		return nil, conn.handleServerUnavailable(packet)
	case steamlang.EMsg_ClientCMList:
		return nil, conn.handleCMList(packet)
	case steamlang.EMsg_ClientSessionToken:
		return nil, conn.handleSessionToken(packet)
	}

	if handler, ok := conn.packetHandlers[packet.MsgType()]; ok {
		return nil, handler(packet)
	}
	return packet, nil
}

func (conn *CMClient) netLoop() error {
	defer log.Println("NetLoop done !")
	for {
		packet, err := conn.RecvPacket()
		if err != nil {
			return err
		}
		if _, err := conn.HandlePacket(packet); err != nil {
			return err
		}
	}
}

func (conn *CMClient) handleMulti(packet *Packet) error {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgMulti]{
		Body: new(steampb.CMsgMulti),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}

	payload := decoder.Body.GetMessageBody()
	if decoder.Body.GetSizeUnzipped() > 0 {
		uncompressed, err := steam.UncompressGzip(payload)
		if err != nil {
			return err
		}
		payload = uncompressed
	}

	rd := bytes.NewReader(payload)
	for rd.Len() > 0 {
		var pktLen uint32
		if err := binary.Read(rd, binary.LittleEndian, &pktLen); err != nil {
			return err
		}
		pktData := make([]byte, pktLen)
		if _, err := io.ReadFull(rd, pktData); err != nil {
			return err
		}

		pkt, err := ParsePacket(pktData)
		if err != nil {
			return err
		}
		if _, err := conn.HandlePacket(pkt); err != nil {
			return err
		}
	}
	return nil
}

func (conn *CMClient) handleServerUnavailable(packet *Packet) error {
	var decoder = &PacketDecoder[*MsgClientServerUnavailable]{
		Body: new(MsgClientServerUnavailable),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}
	return conn.Close()
}

func (conn *CMClient) handleCMList(_ *Packet) error {
	// TODO: Read and update CM list
	return nil
}

func (conn *CMClient) handleSessionToken(packet *Packet) error {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientSessionToken]{
		Body: new(steampb.CMsgClientSessionToken),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}

	// TODO: Set internal session token
	// body.GetToken()
	return nil
}
