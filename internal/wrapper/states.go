package wrapper

import (
	"reflect"

	"github.com/Lucino772/envelop/internal/utils"
)

type ServerState struct {
	Status  ServerState_Status  `json:"status"`
	Players ServerState_Players `json:"players"`
}

type ServerState_Status struct {
	Description string `json:"description"`
}

type ServerState_Players struct {
	Count int                  `json:"count"`
	Max   int                  `json:"max"`
	List  []ServerState_Player `json:"list"`
}

type ServerState_Player struct {
	Id         string         `json:"id"`
	Attributes map[string]any `json:"attributes"`
}

type States struct {
	state       ServerState
	idGenerator func() (string, error)
}

func NewStates(state ServerState) (*States, error) {
	idGenerator, err := utils.NewNanoIDGenerator()
	if err != nil {
		return nil, err
	}
	return &States{state: state, idGenerator: idGenerator}, nil
}

func (s *States) HandleEvent(event Event) (Event, bool) {
	id, err := s.idGenerator()
	if err == nil {
		event.Id = id
	}

	if stateEvent, ok := event.Data.(StateUpdateEvent); ok {
		var updated bool = false
		if !reflect.DeepEqual(s.state, stateEvent.State) {
			updated = true
			s.state = stateEvent.State
		}
		return event, updated
	}
	return event, true
}

func (s *States) Get() ServerState {
	return s.state
}
