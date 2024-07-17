package steamcm

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"golang.org/x/sync/errgroup"
)

type CMConnection struct {
	inner PacketConnection
	errg  *errgroup.Group
}

func NewCMConnection() (*CMConnection, error) {
	s := new(Servers)
	if err := s.Update(); err != nil {
		return nil, err
	}
	server := s.Records()[0]

	conn, err := NewTCPConnection(server.Host, server.Port)
	if err != nil {
		return nil, err
	}
	cm := &CMConnection{
		inner: NewEncryptedConnection(steamlang.EUniverse_Public, conn),
		errg:  &errgroup.Group{},
	}
	cm.errg.Go(cm.netLoop)
	return cm, nil
}

func (conn *CMConnection) SendPacket(packet *Packet) error {
	return conn.inner.SendPacket(packet)
}

func (conn *CMConnection) RecvPacket() (*Packet, error) {
	return conn.inner.RecvPacket()
}

func (conn *CMConnection) Close() error {
	err := conn.inner.Close()
	if err := conn.errg.Wait(); err != nil {
		return err
	}
	return err
}

func (conn *CMConnection) HandlePacket(packet *Packet) (*Packet, error) {
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
	default:
		return nil, nil
	}
}

func (conn *CMConnection) netLoop() error {
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

func (conn *CMConnection) handleMulti(packet *Packet) error {
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

func (conn *CMConnection) handleServerUnavailable(packet *Packet) error {
	var decoder = &PacketDecoder[*MsgClientServerUnavailable]{
		Body: new(MsgClientServerUnavailable),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}
	return conn.Close()
}

func (conn *CMConnection) handleCMList(_ *Packet) error {
	// TODO: Read and update CM list
	return nil
}

func (conn *CMConnection) handleSessionToken(packet *Packet) error {
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
