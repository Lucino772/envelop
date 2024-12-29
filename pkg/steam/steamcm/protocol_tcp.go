package steamcm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
)

const (
	tcpConnectionMagic uint32 = 0x31305456
)

type tcpLayer struct {
	rbuff bytes.Buffer

	outgoingHandler func([]byte) error
	incomingHandler func([]byte) error
}

func NewTCPLayer() *tcpLayer {
	return &tcpLayer{}
}

func (layer *tcpLayer) SetOutgoingHandler(handler func([]byte) error) {
	layer.outgoingHandler = handler
}

func (layer *tcpLayer) SetIncomingHandler(handler func([]byte) error) {
	layer.incomingHandler = handler
}

func (layer *tcpLayer) Send(data []byte) error {
	log.Println("tcp: sending data", len(data))
	var buff bytes.Buffer
	if err := binary.Write(&buff, binary.LittleEndian, uint32(len(data))); err != nil {
		return err
	}
	if err := binary.Write(&buff, binary.LittleEndian, tcpConnectionMagic); err != nil {
		return err
	}
	if _, err := buff.Write(data); err != nil {
		return err
	}

	if layer.outgoingHandler != nil {
		return layer.outgoingHandler(buff.Bytes())
	}
	return nil
}

func (layer *tcpLayer) Handle(data []byte) error {
	log.Println("tcp: handling data", len(data))
	if _, err := layer.rbuff.Write(data); err != nil {
		return err
	}

	for {
		packet, err := layer.tryRead()
		if err != nil {
			return err
		}

		if layer.incomingHandler != nil {
			if err := layer.incomingHandler(packet); err != nil {
				return err
			}
		}
	}
}

func (layer *tcpLayer) tryRead() ([]byte, error) {
	// Check that there is at least enough data
	if layer.rbuff.Len() < 8 {
		return nil, io.ErrUnexpectedEOF
	}

	// Peek at packet len and check that there is enough data
	pktLen := binary.LittleEndian.Uint32(layer.rbuff.Bytes()[:4])
	if layer.rbuff.Len() < int(pktLen)+8 {
		return nil, io.ErrUnexpectedEOF
	}

	// Read the entire packet
	pktData := make([]byte, pktLen+8)
	if _, err := io.ReadFull(&layer.rbuff, pktData); err != nil {
		// If there is still not enough data, consider this as a EOF error. This should
		// not be possible but just in case we make sure the error is different.
		if errors.Is(err, io.ErrUnexpectedEOF) {
			return nil, io.EOF
		}
		return nil, err
	}

	pktMagic := binary.LittleEndian.Uint32(pktData[4:8])
	if pktMagic != tcpConnectionMagic {
		return nil, errors.New("invalid packet magic")
	}

	return pktData[8:], nil
}
