package wrapper

type Event struct {
	Id        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Name      string `json:"name"`
	Data      any    `json:"data"`
}

type ProcessLogEvent struct {
	Value string `json:"value"`
}

func (ev ProcessLogEvent) GetEventName() string {
	return "/process/log"
}

type StateUpdateEvent struct {
	Name string       `json:"name"`
	Data WrapperState `json:"state"`
}

func (ev StateUpdateEvent) GetEventName() string {
	return "/state/update"
}
