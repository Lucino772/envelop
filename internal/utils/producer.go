package utils

import (
	"context"
	"sync"
	"time"
)

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
	producer.mu.Lock()
	defer producer.mu.Unlock()
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
	}
	producer.subscribers = append(producer.subscribers, subscriber)
	return subscriber
}

func (producer *Producer[T]) Close() {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	if producer.closed {
		return
	}

	producer.closed = true
	close(producer.incoming)
	for _, sub := range producer.subscribers[:] {
		sub.Unsubscribe()
	}
}

func (producer *Producer[T]) Run(ctx context.Context) {
	defer producer.Close()
	for {
		select {
		case msg, ok := <-producer.incoming:
			if !ok {
				return
			}
			var wg sync.WaitGroup
			for _, sub := range producer.subscribers {
				wg.Add(1)
				go sub.sendWithTimeout(ctx, msg, &wg)
			}
			wg.Wait()
		case <-ctx.Done():
			return
		}
	}
}

func (sub *subscriber[T]) Messages() <-chan T {
	return sub.messages
}

func (sub *subscriber[T]) Unsubscribe() {
	sub.producer.removeSubscriber(sub)
	close(sub.messages)
}

func (producer *Producer[T]) removeSubscriber(subscriber *subscriber[T]) {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	for ix, sub := range producer.subscribers[:] {
		if sub == subscriber {
			producer.subscribers[ix] = producer.subscribers[len(producer.subscribers)-1]
			producer.subscribers = producer.subscribers[:len(producer.subscribers)-1]
		}
	}
}

func (sub *subscriber[T]) sendWithTimeout(ctx context.Context, message T, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	select {
	case sub.messages <- message:
		return
	case <-ctx.Done():
		return
	}
}
