package wrapper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/Lucino772/envelop/pkg/pubsub"
)

var ErrWrapperContextMissing = errors.New("wrapper context missing")

type wrapperIncomingGrpcKey struct{}

type WrapperContext interface {
	WriteCommand(command string) error
	SendSignal(signal os.Signal) error
	SubscribeLogs() pubsub.Subscriber[string]
	SubscribeEvents() pubsub.Subscriber[Event]
	PublishEvent(event any)
	ReadState(state any) bool
	SubscribeStates() pubsub.Subscriber[any]
	PublishState(state any)
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

func (wp *Wrapper) SubscribeLogs() pubsub.Subscriber[string] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (string, bool) {
		if event, ok := e.Data.(ProcessLogEvent); ok {
			return event.Value, true
		}
		return "", false
	})
}

func (wp *Wrapper) SubscribeEvents() pubsub.Subscriber[Event] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (Event, bool) {
		return e, true
	})
}

func (wp *Wrapper) PublishEvent(event any) {
	wp.eventsProducer.Emit(Event{
		Timestamp: time.Now().Unix(),
		Name:      GetEventName(event),
		Data:      event,
	})
}

func (wp *Wrapper) ReadState(state any) bool {
	if state == nil {
		return false
	}

	value, ok := wp.states[GetStateName(state)]
	if !ok {
		return false
	}

	valuePtr := reflect.ValueOf(value)
	if valuePtr.Kind() != reflect.Ptr {
		return false
	}
	reflect.ValueOf(state).Elem().Set(valuePtr.Elem())
	return true
}

func (wp *Wrapper) SubscribeStates() pubsub.Subscriber[any] {
	return pubsub.NewSubscriber(wp.eventsProducer, func(e Event) (any, bool) {
		if event, ok := e.Data.(StateUpdateEvent); ok {
			return event.Data, true
		}
		return nil, false
	})
}

func (wp *Wrapper) PublishState(state any) {
	wp.PublishEvent(StateUpdateEvent{
		Name: GetStateName(state),
		Data: state,
	})
}
