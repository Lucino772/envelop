package internal

import (
	"context"
	"sync"
	"time"
)

type Producer[T interface{}] struct {
	mu       sync.Mutex
	incoming chan T
	subs     []chan T
	closed   bool
}

func NewProducer[T interface{}]() *Producer[T] {
	return &Producer[T]{
		subs:     make([]chan T, 0),
		incoming: make(chan T, 5),
		closed:   false,
	}
}

func (producer *Producer[T]) Publish(msg T) {
	producer.mu.Lock()
	defer producer.mu.Unlock()
	if producer.closed {
		return
	}
	producer.incoming <- msg
}

func (producer *Producer[T]) Subscribe() <-chan T {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	if producer.closed {
		return nil
	}

	ch := make(chan T)
	producer.subs = append(producer.subs, ch)
	return ch
}

func (producer *Producer[T]) Unsubscribe(channel <-chan T) {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	for ix, sub := range producer.subs[:] {
		if sub == channel {
			producer.subs[ix] = producer.subs[len(producer.subs)-1]
			producer.subs = producer.subs[:len(producer.subs)-1]
			close(sub)
		}
	}
}

func (producer *Producer[T]) Close() {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	if producer.closed {
		return
	}

	producer.closed = true
	close(producer.incoming)
	for _, ch := range producer.subs {
		close(ch)
	}
}

func (producer *Producer[T]) Run() {
	var wg sync.WaitGroup
	for msg := range producer.incoming {
		for _, channel := range producer.subs {
			wg.Add(1)
			producer.publishMessageToChannel(channel, msg, &wg)
		}
	}
	wg.Wait()
}

func (producer *Producer[T]) publishMessageToChannel(channel chan<- T, msg T, wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	defer wg.Done()

	select {
	case channel <- msg:
		return
	case <-ctx.Done():
		return
	}
}
