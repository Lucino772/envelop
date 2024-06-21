package pubsub

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Subscriber[T any] struct {
	messages      chan T
	unsubscribeFn func(*Subscriber[T])
}

func (s *Subscriber[T]) Messages() <-chan T {
	return s.messages
}

func (s *Subscriber[T]) Unsubscribe() {
	s.unsubscribeFn(s)
}

type Producer[T any] struct {
	mu          sync.Mutex
	incoming    chan T
	subscribers []*Subscriber[T]
	closed      bool
}

func NewProducer[T any](chanSize int) *Producer[T] {
	return &Producer[T]{
		subscribers: make([]*Subscriber[T], 0),
		incoming:    make(chan T, chanSize),
		closed:      false,
	}
}

func (p *Producer[T]) Publish(v T) {
	if !p.closed {
		p.incoming <- v
	}
}

func (p *Producer[T]) Subscribe() *Subscriber[T] {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil
	}
	sub := &Subscriber[T]{
		messages:      make(chan T),
		unsubscribeFn: p.unsubscribe,
	}
	p.subscribers = append(p.subscribers, sub)
	return sub
}

func (p *Producer[T]) Close() {
	if p.closed {
		return
	}
	p.closed = true
	close(p.incoming)
	for _, sub := range p.subscribers[:] {
		p.unsubscribe(sub)
	}
}

func (p *Producer[T]) Run(ctx context.Context) error {
	defer p.Close()

	for {
		select {
		case msg, ok := <-p.incoming:
			if !ok {
				return nil
			}
			eg := new(errgroup.Group)
			for _, sub := range p.subscribers {
				eg.Go(p.messageSender(ctx, msg, sub))
			}
			if err := eg.Wait(); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *Producer[T]) messageSender(parent context.Context, msg T, s *Subscriber[T]) func() error {
	return func() error {
		ctx, cancel := context.WithTimeout(parent, 5*time.Second)
		defer cancel()

		select {
		case s.messages <- msg:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *Producer[T]) unsubscribe(s *Subscriber[T]) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for ix, sub := range p.subscribers[:] {
		if sub == s {
			p.subscribers[ix] = p.subscribers[len(p.subscribers)-1]
			p.subscribers = p.subscribers[:len(p.subscribers)-1]
			close(sub.messages)
		}
	}
}
