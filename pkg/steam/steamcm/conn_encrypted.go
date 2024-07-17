package steamcm

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"hash/crc32"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
)

const (
	state_Disconnected int = 0
	state_Connected    int = 1
	state_Challenged   int = 2
	state_Encrypted    int = 3
)

type EncryptedConnection struct {
	Universe steamlang.EUniverse

	inner     Connection
	state     int
	encrypter Encrypter
}

func NewEncryptedConnection(universe steamlang.EUniverse, conn Connection) PacketConnection {
	return &EncryptedConnection{
		Universe:  universe,
		inner:     conn,
		state:     state_Disconnected,
		encrypter: nil,
	}
}

func (conn *EncryptedConnection) WritePacket(packet []byte) error {
	if conn.state == state_Encrypted {
		encrypted, err := conn.encrypter.Encrypt(packet)
		if err != nil {
			return err
		}
		packet = encrypted
	}
	return conn.inner.WritePacket(packet)
}

func (conn *EncryptedConnection) ReadPacket() ([]byte, error) {
	packet, err := conn.inner.ReadPacket()
	if err != nil {
		return nil, err
	}

	if conn.state == state_Encrypted {
		decrypted, err := conn.encrypter.Decrypt(packet)
		if err != nil {
			return nil, err
		}
		packet = decrypted
	}

	return packet, nil
}

func (conn *EncryptedConnection) Close() error {
	return conn.inner.Close()
}

func (conn *EncryptedConnection) SendPacket(packet *Packet) error {
	return conn.WritePacket(packet.Bytes())
}

func (conn *EncryptedConnection) RecvPacket() (*Packet, error) {
	data, err := conn.ReadPacket()
	if err != nil {
		return nil, err
	}
	return ParsePacket(data)
}

func (conn *EncryptedConnection) HandlePacket(packet *Packet) (*Packet, error) {
	if conn.state == state_Encrypted {
		return packet, nil
	}

	switch packet.MsgType() {
	case steamlang.EMsg_ChannelEncryptRequest:
		return nil, conn.handleEncryptRequest(packet)
	case steamlang.EMsg_ChannelEncryptResult:
		return nil, conn.handleEncryptResult(packet)
	default:
		return nil, nil
	}
}

func (conn *EncryptedConnection) handleEncryptRequest(packet *Packet) error {
	var decoder = &PacketDecoder[*MsgChannelEncryptRequest]{
		Body: new(MsgChannelEncryptRequest),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}
	randomChallenge := decoder.Payload

	if decoder.Body.ProtoVersion != 1 {
		return errors.New("version mismatch")
	}
	if decoder.Body.Universe != conn.Universe {
		return errors.New("unexpected universe")
	}

	pubKey := steam.GetUniversePublicKey(decoder.Body.Universe)
	if pubKey == nil {
		return errors.New("invalid universe")
	}

	var dataToEncrypt bytes.Buffer
	tempSessionKey := make([]byte, 32)
	if _, err := rand.Read(tempSessionKey); err != nil {
		return err
	}
	if _, err := dataToEncrypt.Write(tempSessionKey); err != nil {
		return err
	}
	if len(randomChallenge) > 0 {
		if _, err := dataToEncrypt.Write(randomChallenge); err != nil {
			return err
		}
		conn.encrypter = NewEncrypter(tempSessionKey, tempSessionKey[:16])
	} else {
		conn.encrypter = NewEncrypter(tempSessionKey, nil)
	}

	encryptedData, err := steam.EncryptOAEPSha1(pubKey, dataToEncrypt.Bytes())
	if err != nil {
		return err
	}
	keyCrc := crc32.ChecksumIEEE(encryptedData)

	return conn.sendEncryptResponse(decoder.Body.ProtoVersion, encryptedData, keyCrc)
}

func (conn *EncryptedConnection) sendEncryptResponse(version uint32, challengeData []byte, crc uint32) error {
	encoder := NewPacketEncoder(steamlang.EMsg_ChannelEncryptResponse)
	encoder.Body = &MsgChannelEncryptResponse{
		ProtoVersion: version,
		KeySize:      128,
	}
	if _, err := encoder.Data.Write(challengeData); err != nil {
		return err
	}
	if err := binary.Write(encoder.Data, binary.LittleEndian, crc); err != nil {
		return err
	}
	if err := binary.Write(encoder.Data, binary.LittleEndian, uint32(0)); err != nil {
		return err
	}
	packet, err := encoder.Encode()
	if err != nil {
		return err
	}
	if err := conn.SendPacket(packet); err != nil {
		return err
	}
	conn.state = state_Challenged
	return nil
}

func (conn *EncryptedConnection) handleEncryptResult(packet *Packet) error {
	var decoder = &PacketDecoder[*MsgChannelEncryptResult]{
		Body: new(MsgChannelEncryptResult),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}

	// FIXME: What should we do if result is not ok, disconnect ?
	if decoder.Body.Result == steamlang.EResult_OK {
		conn.state = state_Encrypted
	}
	return nil
}
