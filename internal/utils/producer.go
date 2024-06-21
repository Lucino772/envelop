package utils

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Subscriber[T any] interface {
	Messages() <-chan T
	Unsubscribe()
}

type subscriber[T interface{}] struct {
	messages chan T
	producer *Producer[T]
}

type Producer[T interface{}] struct {
	mu          sync.Mutex
	incoming    chan T
	subscribers []*subscriber[T]
	closed      bool
}

func NewProducer[T interface{}]() *Producer[T] {
	return &Producer[T]{
		subscribers: make([]*subscriber[T], 0),
		incoming:    make(chan T, 5),
		closed:      false,
	}
}

func (producer *Producer[T]) Publish(msg T) {
	if !producer.closed {
		producer.incoming <- msg
	}
}

func (producer *Producer[T]) Subscribe() *subscriber[T] {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	if producer.closed {
		return nil
	}

	subscriber := &subscriber[T]{
		messages: make(chan T),
		producer: producer,
	}
	producer.subscribers = append(producer.subscribers, subscriber)
	return subscriber
}

func (producer *Producer[T]) Close() {
	if producer.closed {
		return
	}

	producer.closed = true
	close(producer.incoming)
	for _, sub := range producer.subscribers[:] {
		sub.Unsubscribe()
	}
}

func (producer *Producer[T]) Run(ctx context.Context) error {
	defer producer.Close()

	for {
		select {
		case msg, ok := <-producer.incoming:
			if !ok {
				return nil
			}
			g, newCtx := errgroup.WithContext(ctx)
			for _, sub := range producer.subscribers {
				sub := sub
				g.Go(func() error {
					return sub.sendWithTimeout(newCtx, msg)
				})
			}
			if err := g.Wait(); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (sub *subscriber[T]) Messages() <-chan T {
	return sub.messages
}

func (sub *subscriber[T]) Unsubscribe() {
	sub.producer.removeSubscriber(sub)
}

func (producer *Producer[T]) removeSubscriber(subscriber *subscriber[T]) {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	for ix, sub := range producer.subscribers[:] {
		if sub == subscriber {
			producer.subscribers[ix] = producer.subscribers[len(producer.subscribers)-1]
			producer.subscribers = producer.subscribers[:len(producer.subscribers)-1]
			close(sub.messages)
		}
	}
}

func (sub *subscriber[T]) sendWithTimeout(ctx context.Context, message T) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	select {
	case sub.messages <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
