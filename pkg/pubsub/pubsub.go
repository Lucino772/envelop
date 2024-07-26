package pubsub

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Subscriber[T any] interface {
	Receive() <-chan T
	Send() chan<- T
	Unsubscribe()
}

type Publisher[T any] interface {
	Publish(v T)
	Subscribe() Subscriber[T]
	Run(ctx context.Context) error
	Stop()
}

type subscriber[T any] struct {
	channel       chan T
	unsubscribeFn func(Subscriber[T])
	closed        bool
}

func NewSubscriber[T any](unsubscribeFn func(Subscriber[T])) Subscriber[T] {
	return &subscriber[T]{
		channel:       make(chan T),
		unsubscribeFn: unsubscribeFn,
		closed:        false,
	}
}

func (s *subscriber[T]) Receive() <-chan T {
	return s.channel
}

func (s *subscriber[T]) Send() chan<- T {
	return s.channel
}

func (s *subscriber[T]) Unsubscribe() {
	if !s.closed {
		if s.unsubscribeFn != nil {
			s.unsubscribeFn(s)
		}
		close(s.channel)
		s.closed = true
	}
}

type publisher[T any] struct {
	mu          sync.Mutex
	incoming    chan T
	subscribers []Subscriber[T]
	processMsg  func(T) (T, bool)
	closed      bool
}

func NewPublisher[T any](chanSize int, processMsg func(T) (T, bool)) Publisher[T] {
	return &publisher[T]{
		subscribers: make([]Subscriber[T], 0),
		incoming:    make(chan T, chanSize),
		processMsg:  processMsg,
		closed:      false,
	}
}

func (p *publisher[T]) Publish(v T) {
	if !p.closed {
		p.incoming <- v
	}
}

func (p *publisher[T]) Subscribe() Subscriber[T] {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil
	}
	sub := NewSubscriber(p.unsubscribe)
	p.subscribers = append(p.subscribers, sub)
	return sub
}

func (p *publisher[T]) Stop() {
	if !p.closed {
		p.closed = true
		close(p.incoming)
	}
}

func (p *publisher[T]) Run(ctx context.Context) error {
	defer func() {
		for _, sub := range p.subscribers[:] {
			sub.Unsubscribe()
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

func (p *publisher[T]) getSender(ctx context.Context, msg T, s Subscriber[T]) func() error {
	return func() error {
		select {
		case s.Send() <- msg:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *publisher[T]) unsubscribe(s Subscriber[T]) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for ix, sub := range p.subscribers[:] {
		if sub == s {
			p.subscribers[ix] = p.subscribers[len(p.subscribers)-1]
			p.subscribers = p.subscribers[:len(p.subscribers)-1]
		}
	}
}
