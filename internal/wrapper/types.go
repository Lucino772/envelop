package wrapper

import (
	"context"

	"google.golang.org/grpc"
)

type WrapperTask func(context.Context)

type WrapperService interface {
	Register(grpc.ServiceRegistrar)
}

type WrapperRegistrar interface {
	AddService(WrapperService)
	AddTask(WrapperTask)
}

type WrapperModule interface {
	Register(WrapperRegistrar)
}

type WrapperEvent interface {
	GetEventName() string
}

type WrapperState interface {
	GetStateName() string
	Equals(WrapperState) bool
}

type WrapperStateAccessor[T WrapperState] interface {
	Get() T
	Set(T)
}

type WrapperSubscriber[T interface{}] interface {
	Messages() <-chan T
	Unsubscribe()
}
