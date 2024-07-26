package wrapper

import (
	"context"
	"reflect"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type StateSubscriber struct {
	states        chan WrapperState
	unsubscribeFn func(*StateSubscriber)
}

func (s *StateSubscriber) Messages() <-chan WrapperState {
	return s.states
}

func (s *StateSubscriber) Unsubscribe() {
	s.unsubscribeFn(s)
}

type StateManager struct {
	mu          sync.Mutex
	incoming    chan WrapperState
	states      map[string]WrapperState
	subscribers []*StateSubscriber
	closed      bool
}

func NewStateManager(chanSize int) *StateManager {
	manager := &StateManager{
		incoming:    make(chan WrapperState, chanSize),
		states:      make(map[string]WrapperState),
		subscribers: make([]*StateSubscriber, 0),
		closed:      false,
	}
	manager.setState(&ProcessStatusState{
		Description: "Unknown",
	})
	manager.setState(&PlayerState{
		Count:   0,
		Max:     0,
		Players: []string{},
	})
	return manager
}

func (manager *StateManager) Publish(v WrapperState) {
	if !manager.closed {
		manager.incoming <- v
	}
}

func (manager *StateManager) Subscribe() *StateSubscriber {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if manager.closed {
		return nil
	}
	sub := &StateSubscriber{
		states:        make(chan WrapperState),
		unsubscribeFn: manager.unsubscribe,
	}
	manager.subscribers = append(manager.subscribers, sub)
	return sub
}

func (manager *StateManager) Read(state WrapperState) bool {
	if state == nil {
		return false
	}

	value, ok := manager.states[state.GetStateName()]
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

func (manager *StateManager) Close() {
	if manager.closed {
		return
	}
	manager.closed = true
	close(manager.incoming)
	for _, sub := range manager.subscribers[:] {
		manager.unsubscribe(sub)
	}
}

func (manager *StateManager) Run(ctx context.Context) error {
	defer manager.Close()

	for {
		select {
		case state, ok := <-manager.incoming:
			if !ok {
				return nil
			}
			if manager.setState(state) {
				eg := new(errgroup.Group)
				for _, sub := range manager.subscribers {
					eg.Go(manager.messageSender(ctx, state, sub))
				}
				if err := eg.Wait(); err != nil {
					return err
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (manager *StateManager) setState(state WrapperState) bool {
	currentState, ok := manager.states[state.GetStateName()]

	var updated bool = false
	if !ok {
		manager.states[state.GetStateName()] = state
		updated = true
	} else if !currentState.Equals(state) {
		manager.states[state.GetStateName()] = state
		updated = true
	}
	return updated
}

func (manager *StateManager) messageSender(parent context.Context, state WrapperState, s *StateSubscriber) func() error {
	return func() error {
		ctx, cancel := context.WithTimeout(parent, 5*time.Second)
		defer cancel()

		select {
		case s.states <- state:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (manager *StateManager) unsubscribe(s *StateSubscriber) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	for ix, sub := range manager.subscribers[:] {
		if sub == s {
			manager.subscribers[ix] = manager.subscribers[len(manager.subscribers)-1]
			manager.subscribers = manager.subscribers[:len(manager.subscribers)-1]
			close(sub.states)
		}
	}
}
