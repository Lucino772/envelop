package steamcm

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"log"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
)

const (
	state_Disconnected int = 0
	state_Connected    int = 1
	state_Challenged   int = 2
	state_Encrypted    int = 3
)

type encryptedLayer struct {
	Universe steamlang.EUniverse

	state           int
	encrypter       Encrypter
	outgoingHandler func([]byte) error
	incomingHandler func([]byte) error
}

func NewEncryptedLayer(universe steamlang.EUniverse) *encryptedLayer {
	return &encryptedLayer{
		Universe:  universe,
		state:     state_Disconnected,
		encrypter: nil,
	}
}

func (layer *encryptedLayer) SetOutgoingHandler(handler func([]byte) error) {
	layer.outgoingHandler = handler
}

func (layer *encryptedLayer) SetIncomingHandler(handler func([]byte) error) {
	layer.incomingHandler = handler
}

func (layer *encryptedLayer) Send(data []byte) error {
	log.Println("encrypted: sending data", len(data))

	if layer.state == state_Encrypted {
		encrypted, err := layer.encrypter.Encrypt(data)
		if err != nil {
			return err
		}
		data = encrypted
	}

	if layer.outgoingHandler != nil {
		layer.outgoingHandler(data)
	}
	return nil
}

func (layer *encryptedLayer) Handle(data []byte) error {
	log.Println("encrypted: handle data", len(data))

	if layer.state == state_Encrypted {
		decrypted, err := layer.encrypter.Decrypt(data)
		if err != nil {
			return err
		}
		if layer.incomingHandler != nil {
			return layer.incomingHandler(decrypted)
		}
		return nil
	}

	packet, err := ParsePacket(data)
	if err != nil {
		return err
	}

	switch packet.MsgType() {
	case steamlang.EMsg_ChannelEncryptRequest:
		return layer.handleEncryptRequest(packet)
	case steamlang.EMsg_ChannelEncryptResult:
		return layer.handleEncryptResult(packet)
	default:
		return nil
	}
}

func (layer *encryptedLayer) handleEncryptRequest(packet *Packet) error {
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
	if decoder.Body.Universe != layer.Universe {
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
		layer.encrypter = NewEncrypter(tempSessionKey, tempSessionKey[:16])
	} else {
		layer.encrypter = NewEncrypter(tempSessionKey, nil)
	}

	encryptedData, err := steam.EncryptOAEPSha1(pubKey, dataToEncrypt.Bytes())
	if err != nil {
		return err
	}
	keyCrc := crc32.ChecksumIEEE(encryptedData)

	return layer.sendEncryptResponse(decoder.Body.ProtoVersion, encryptedData, keyCrc)
}

func (layer *encryptedLayer) sendEncryptResponse(version uint32, challengeData []byte, crc uint32) error {
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
	layer.state = state_Challenged
	return layer.Send(packet.Bytes())
}

func (layer *encryptedLayer) handleEncryptResult(packet *Packet) error {
	var decoder = &PacketDecoder[*MsgChannelEncryptResult]{
		Body: new(MsgChannelEncryptResult),
	}
	if err := decoder.Decode(packet); err != nil {
		return err
	}

	// FIXME: What should we do if result is not ok, disconnect ?
	if decoder.Body.Result == steamlang.EResult_OK {
		layer.state = state_Encrypted
	}
	return nil
}
