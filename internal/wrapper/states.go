package wrapper

func GetStateName(state any) string {
	switch state.(type) {
	case ProcessStatusState, *ProcessStatusState:
		return "/process/status"
	case PlayerState, *PlayerState:
		return "/player/list"
	default:
		return "/unknown"
	}
}

type ProcessStatusState struct {
	Description string `json:"description"`
}

type PlayerState struct {
	Count   int      `json:"count"`
	Max     int      `json:"max"`
	Players []string `json:"players"`
}
