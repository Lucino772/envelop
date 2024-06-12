package wrapper

import (
	"sync"
	"time"

	"github.com/Lucino772/envelop/internal/utils"
)

type ProcessStatusState struct {
	Description string `json:"description"`
}

func (state ProcessStatusState) GetStateName() string {
	return "/process/status"
}

type wrapperStateAccessor[T WrapperState] struct {
	eventsProducer *utils.Producer[Event]
	stateObj       T
	mu             sync.Mutex
}

func NewWrapperStateAccessor[T WrapperState](eventsProducer *utils.Producer[Event], initialState T) *wrapperStateAccessor[T] {
	return &wrapperStateAccessor[T]{
		eventsProducer: eventsProducer,
		stateObj:       initialState,
	}
}

func (wsa *wrapperStateAccessor[T]) Get() T {
	return wsa.stateObj
}

func (wsa *wrapperStateAccessor[T]) Set(state T) {
	wsa.mu.Lock()
	defer wsa.mu.Unlock()
	wsa.stateObj = state
	event := StateUpdateEvent{
		Name: state.GetStateName(),
		Data: state,
	}
	wsa.eventsProducer.Publish(Event{
		Id:        "", // TODO: Get unique ID
		Name:      event.GetEventName(),
		Timestamp: time.Now().Unix(),
		Data:      event,
	})
}
