package wrapper

import (
	"context"

	"github.com/Lucino772/envelop/internal"
	pb "github.com/Lucino772/envelop/pkg/protobufs"
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

type Wrapper interface {
	WriteCommand(string)
	GetLogsProducer() *internal.Producer[string]
	GetEventsProducer() *internal.Producer[*pb.Event]
}
