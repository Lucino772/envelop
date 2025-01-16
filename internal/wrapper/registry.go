package wrapper

import "log/slog"

type Registry struct {
	Tasks    []Task
	Services []Service

	Stoppers        map[string]func(map[string]any) Stopper
	LoggingHandlers map[string]func(map[string]any) slog.Handler
}

func NewRegistry() *Registry {
	return &Registry{
		Tasks:           make([]Task, 0),
		Services:        make([]Service, 0),
		Stoppers:        make(map[string]func(map[string]any) Stopper),
		LoggingHandlers: make(map[string]func(map[string]any) slog.Handler),
	}
}
