package wrapper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Lucino772/envelop/internal/utils"
)

var ErrWrapperContextMissing = errors.New("wrapper context missing")

type wrapperIncomingGrpcKey struct{}

type WrapperEvent interface {
	GetEventName() string
}

type WrapperState interface {
	GetStateName() string
	Equals(WrapperState) bool
}

type WrapperStateProperty[T WrapperState] interface {
	Get() T
	Set(T)
}

type WrapperContext interface {
	WriteCommand(command string) error
	SendSignal(signal os.Signal) error
	SubscribeLogs() utils.Subscriber[string]
	SubscribeEvents() utils.Subscriber[Event]
	PublishEvent(event WrapperEvent)
	ProcessStatusState() WrapperStateProperty[ProcessStatusState]
	PlayerState() WrapperStateProperty[PlayerState]
}

func FromContext(ctx context.Context) (WrapperContext, error) {
	wrapper, ok := ctx.Value(wrapperIncomingGrpcKey{}).(WrapperContext)
	if !ok {
		return nil, ErrWrapperContextMissing
	}
	return wrapper, nil
}

func (wp *Wrapper) withContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, wrapperIncomingGrpcKey{}, wp)
}

func (wp *Wrapper) WriteCommand(command string) error {
	_, err := wp.stdinWriter.Write([]byte(fmt.Sprintf("%s\n", command)))
	return err
}

func (wp *Wrapper) SendSignal(signal os.Signal) error {
	process, err := os.FindProcess(wp.cmd.Status().PID)
	if err != nil {
		return err
	}
	return process.Signal(signal)
}

func (wp *Wrapper) SubscribeLogs() utils.Subscriber[string] {
	return wp.logsProducer.Subscribe()
}

func (wp *Wrapper) SubscribeEvents() utils.Subscriber[Event] {
	return wp.eventsProducer.Subscribe()
}

func (wp *Wrapper) PublishEvent(event WrapperEvent) {
	wp.eventsProducer.Publish(Event{
		Id:        "", // TODO: Get Unique ID
		Timestamp: time.Now().Unix(),
		Name:      event.GetEventName(),
		Data:      event,
	})
}

func (wp *Wrapper) ProcessStatusState() WrapperStateProperty[ProcessStatusState] {
	return wp.processStatusState
}

func (wp *Wrapper) PlayerState() WrapperStateProperty[PlayerState] {
	return wp.playerState
}
