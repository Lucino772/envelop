package pubsub

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrProducerClosed error = errors.New("producer closed")
)

type producer[T any] struct {
	mu          sync.Mutex
	incoming    chan T
	subscribers []Sender[T]
	processMsg  func(T) (T, bool)
	closed      bool
}

func NewProducer[T any](size int, processMsg func(T) (T, bool)) Producer[T] {
	return &producer[T]{
		subscribers: make([]Sender[T], 0),
		incoming:    make(chan T, size),
		processMsg:  processMsg,
		closed:      false,
	}
}

func (p *producer[T]) Emit(v T) {
	if !p.closed {
		p.incoming <- v
	}
}

func (p *producer[T]) Attach(s Sender[T]) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrProducerClosed
	}
	p.subscribers = append(p.subscribers, s)
	return nil
}

func (p *producer[T]) Detach(s Sender[T]) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for ix, sub := range p.subscribers[:] {
		if sub == s {
			p.subscribers[ix] = p.subscribers[len(p.subscribers)-1]
			p.subscribers = p.subscribers[:len(p.subscribers)-1]
		}
	}
}

func (p *producer[T]) Stop() {
	if !p.closed {
		p.closed = true
		close(p.incoming)
	}
}

func (p *producer[T]) Run(ctx context.Context) error {
	defer func() {
		for _, sub := range p.subscribers[:] {
			sub.Close()
		}
	}()

	for {
		select {
		case msg, ok := <-p.incoming:
			if !ok {
				return nil
			}

			var doNotify bool = true
			if p.processMsg != nil {
				msg, doNotify = p.processMsg(msg)
			}

			if doNotify {
				wg, _ := errgroup.WithContext(ctx)
				for _, sub := range p.subscribers {
					wg.Go(p.getSender(ctx, msg, sub))
				}
				if err := wg.Wait(); err != nil {
					return err
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *producer[T]) getSender(ctx context.Context, msg T, s Sender[T]) func() error {
	return func() error {
		return s.Send(ctx, msg)
	}
}
