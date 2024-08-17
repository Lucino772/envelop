package wrapper

import (
	"time"
)

func GetEventName(event any) string {
	switch event.(type) {
	case LogEvent, *LogEvent:
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

type LogEvent struct {
	Time    time.Time      `json:"timestamp"`
	Message string         `json:"message"`
	Level   string         `json:"level"`
	Data    map[string]any `json:"data"`
}

type StateUpdateEvent struct {
	Name string `json:"name"`
	Data any    `json:"state"`
}
