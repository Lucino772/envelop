package wrapper

import (
	"context"
	"net"
	"time"

	"github.com/Lucino772/envelop/internal"
	"google.golang.org/grpc"
)

type DefaultWrapper struct {
	ProcessStatusState WrapperStateAccessor[ProcessStatusState]

	process        *WrapperProcess
	logsProducer   *internal.Producer[string]
	eventsProducer *internal.Producer[Event]
	services       []WrapperService
	tasks          []WrapperTask
}

func NewDefaultWrapper(process *WrapperProcess, logsProducer *internal.Producer[string], eventsProducer *internal.Producer[Event]) *DefaultWrapper {
	return &DefaultWrapper{
		ProcessStatusState: NewWrapperStateAccessor(eventsProducer, ProcessStatusState{
			Description: "Unknown",
		}),

		process:        process,
		logsProducer:   logsProducer,
		eventsProducer: eventsProducer,
		services:       make([]WrapperService, 0),
		tasks:          make([]WrapperTask, 0),
	}
}

func (wp *DefaultWrapper) AddService(service WrapperService) {
	wp.services = append(wp.services, service)
}

func (wp *DefaultWrapper) AddTask(task WrapperTask) {
	wp.tasks = append(wp.tasks, task)
}

func (wp *DefaultWrapper) WriteCommand(command string) error {
	_, err := wp.process.Write(command)
	return err
}

func (wp *DefaultWrapper) PublishEvent(event WrapperEvent) {
	wp.eventsProducer.Publish(Event{
		Id:        "", // TODO: Get Unique ID
		Timestamp: time.Now().Unix(),
		Name:      event.GetEventName(),
		Data:      event,
	})
}

func (wp *DefaultWrapper) GetLogsProducer() *internal.Producer[string] {
	return wp.logsProducer
}

func (wp *DefaultWrapper) GetEventsProducer() *internal.Producer[Event] {
	return wp.eventsProducer
}

func (wp *DefaultWrapper) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	wp.tasks = append(wp.tasks, wp.logsProducer, wp.eventsProducer)
	grpcServer, err := wp.startGrpc()
	if err != nil {
		return err
	}
	defer grpcServer.Stop()
	for _, task := range wp.tasks {
		go task.Run(ctx)
	}
	wp.process.Run(ctx)
	return nil
}

func (wp *DefaultWrapper) startGrpc() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:8791")
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(getGrpcUnaryInterceptor(wp)),
		grpc.StreamInterceptor(getGrpcStreamInterceptor(wp)),
	)
	for _, service := range wp.services {
		service.Register(grpcServer)
	}
	go grpcServer.Serve(lis)
	return grpcServer, nil
}
