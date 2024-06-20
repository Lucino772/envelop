package wrapper

import "slices"

type ProcessStatusState struct {
	Description string `json:"description"`
}

func (state ProcessStatusState) GetStateName() string {
	return "/process/status"
}

func (state ProcessStatusState) Equals(otherState WrapperState) bool {
	if val, ok := otherState.(ProcessStatusState); ok {
		return state.Description == val.Description
	}
	return false
}

type PlayerState struct {
	Count   int      `json:"count"`
	Max     int      `json:"max"`
	Players []string `json:"players"`
}

func (state PlayerState) GetStateName() string {
	return "/player/list"
}

func (state PlayerState) Equals(otherState WrapperState) bool {
	if val, ok := otherState.(PlayerState); ok {
		return state.Count == val.Count &&
			state.Max == val.Max &&
			slices.Equal(state.Players, val.Players)
	}
	return false
}
