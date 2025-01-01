package steamcm

import (
	"bytes"
	"sync"
)

type buffer struct {
	buff bytes.Buffer
	mu   sync.Mutex
	cond *sync.Cond
}

func NewBuffer() *buffer {
	buff := &buffer{}
	buff.cond = sync.NewCond(&buff.mu)
	return buff
}

func (b *buffer) Read(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.buff.Len() == 0 {
		b.cond.Wait()
	}
	return b.buff.Read(p)
}

func (b *buffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	n, err := b.buff.Write(p)
	b.cond.Signal()
	return n, err
}
