package wrapper

func GetEventName(event any) string {
	switch event.(type) {
	case ProcessLogEvent, *ProcessLogEvent:
		return "/process/log"
	case StateUpdateEvent, *StateUpdateEvent:
		return "/state/update"
	default:
		return "/unkown"
	}
}

type Event struct {
	Id        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Name      string `json:"name"`
	Data      any    `json:"data"`
}

type ProcessLogEvent struct {
	Message string `json:"message"`
}

type StateUpdateEvent struct {
	State ServerState `json:"state"`
}
