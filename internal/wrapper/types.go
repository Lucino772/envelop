package wrapper

import (
	"context"

	"github.com/Lucino772/envelop/internal"
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

type Wrapper interface {
	WriteCommand(string) error
	PublishEvent(WrapperEvent)
	GetLogsProducer() *internal.Producer[string]
	GetEventsProducer() *internal.Producer[Event]

	// States
	GetProcessStatusState() WrapperStateAccessor[ProcessStatusState]
}
