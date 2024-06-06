package wrapper

import (
	"context"
	"net"

	"github.com/Lucino772/envelop/internal"
	"google.golang.org/grpc"
)

type DefaultWrapper struct {
	process      *WrapperProcess
	logsProducer *internal.Producer[string]
	services     []WrapperService
	tasks        []WrapperTask
}

func NewDefaultWrapper(process *WrapperProcess, logsProducer *internal.Producer[string]) *DefaultWrapper {
	return &DefaultWrapper{
		process:      process,
		logsProducer: logsProducer,
		services:     make([]WrapperService, 0),
		tasks:        make([]WrapperTask, 0),
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

func (wp *DefaultWrapper) GetLogsProducer() *internal.Producer[string] {
	return wp.logsProducer
}

func (wp *DefaultWrapper) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	wp.tasks = append(wp.tasks, wp.logsProducer)
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
	grpcServer := grpc.NewServer()
	for _, service := range wp.services {
		service.Register(grpcServer)
	}
	go grpcServer.Serve(lis)
	return grpcServer, nil
}
