package wrapper

import (
	"reflect"

	"github.com/Lucino772/envelop/pkg/pubsub"
)

type StatePublisher struct {
	pubsub.Publisher[WrapperState]
	states map[string]WrapperState
}

func NewStatePublisher(chanSize int) *StatePublisher {
	states := make(map[string]WrapperState, 0)
	publisher := &StatePublisher{
		Publisher: pubsub.NewPublisher(chanSize, getStateMsgProcessor(states)),
		states:    states,
	}
	publisher.setState(&ProcessStatusState{
		Description: "Unknown",
	})
	publisher.setState(&PlayerState{
		Count:   0,
		Max:     0,
		Players: []string{},
	})
	return publisher
}

func getStateMsgProcessor(states map[string]WrapperState) func(WrapperState) (WrapperState, bool) {
	return func(state WrapperState) (WrapperState, bool) {
		currentState, ok := states[state.GetStateName()]

		var updated bool = false
		if !ok {
			states[state.GetStateName()] = state
			updated = true
		} else if !currentState.Equals(state) {
			states[state.GetStateName()] = state
			updated = true
		}
		return state, updated
	}
}

func (publisher *StatePublisher) Read(state WrapperState) bool {
	if state == nil {
		return false
	}

	value, ok := publisher.states[state.GetStateName()]
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

func (publisher *StatePublisher) setState(state WrapperState) bool {
	currentState, ok := publisher.states[state.GetStateName()]

	var updated bool = false
	if !ok {
		publisher.states[state.GetStateName()] = state
		updated = true
	} else if !currentState.Equals(state) {
		publisher.states[state.GetStateName()] = state
		updated = true
	}
	return updated
}
