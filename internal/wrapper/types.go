package wrapper

import (
	"context"

	"google.golang.org/grpc"
)

type WrapperTask interface {
	Run(context.Context)
}

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
}

type WrapperStateAccessor[T WrapperState] interface {
	Get() T
	Set(T)
}

type WrapperSubscriber[T interface{}] interface {
	Messages() <-chan T
	Unsubscribe()
}

type Wrapper interface {
	WriteCommand(string) error
	SubscribeLogs() WrapperSubscriber[string]
	SubscribeEvents() WrapperSubscriber[Event]
	PublishEvent(WrapperEvent)

	// States
	GetProcessStatusState() WrapperStateAccessor[ProcessStatusState]
}
