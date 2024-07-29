package pubsub

import "context"

type subscriber[T any, K any] struct {
	producer        Producer[K]
	incoming        chan T
	closed          bool
	processIncoming func(K) (T, bool)
}

func NewSubscriber[T any, K any](p Producer[K], processIncoming func(K) (T, bool)) Subscriber[T] {
	sub := &subscriber[T, K]{
		producer:        p,
		incoming:        make(chan T),
		closed:          false,
		processIncoming: processIncoming,
	}
	if err := p.Attach(sub); err != nil {
		return nil
	}
	return sub
}

func (s *subscriber[T, K]) Send(ctx context.Context, val K) error {
	if value, ok := s.processIncoming(val); ok {
		select {
		case s.incoming <- value:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func (s *subscriber[T, K]) Close() {
	if !s.closed {
		s.producer.Detach(s)
		close(s.incoming)
		s.closed = true
	}
}

func (s *subscriber[T, K]) Receive() <-chan T {
	return s.incoming
}
