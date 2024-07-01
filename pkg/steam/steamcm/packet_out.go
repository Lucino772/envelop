package steamcm

import (
	"bytes"
	"io"
)

type PacketWriter struct {
	header io.WriterTo
	body   io.WriterTo
	data   bytes.Buffer
}

func (w *PacketWriter) WithHeader(h io.WriterTo) {
	w.header = h
}

func (w *PacketWriter) WithBody(b io.WriterTo) {
	w.body = b
}

func (w *PacketWriter) Write(p []byte) (int, error) {
	return w.data.Write(p)
}

func (w *PacketWriter) Bytes() []byte {
	var buf bytes.Buffer
	if _, err := w.header.WriteTo(&buf); err != nil {
		return nil
	}
	if _, err := w.body.WriteTo(&buf); err != nil {
		return nil
	}
	if _, err := w.data.WriteTo(&buf); err != nil {
		return nil
	}
	return buf.Bytes()
}
