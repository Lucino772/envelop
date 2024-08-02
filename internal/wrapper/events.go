package wrapper

func GetEventName(event any) string {
	switch event.(type) {
	case ProcessLogEvent:
		return "/process/log"
	case StateUpdateEvent:
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
	Value string `json:"value"`
}

type StateUpdateEvent struct {
	Name string       `json:"name"`
	Data WrapperState `json:"state"`
}
