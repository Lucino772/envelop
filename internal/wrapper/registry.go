package wrapper

import "log/slog"

type Registry struct {
	Tasks    []Task
	Services []Service

	Stoppers        map[string]func(map[string]any) Stopper
	LoggingHandlers map[string]func(map[string]any) slog.Handler
}
