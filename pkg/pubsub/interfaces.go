package pubsub

import "context"

type Producer[T any] interface {
	Emit(v T)
	Attach(Sender[T]) error
	Detach(Sender[T])
	Run(ctx context.Context) error
	Stop()
}

type Sender[T any] interface {
	Send(ctx context.Context, v T) error
	Close()
}

type Subscriber[T any] interface {
	Receive() <-chan T
	Close()
}
