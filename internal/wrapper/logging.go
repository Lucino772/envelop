package wrapper

import (
	"context"
	"log/slog"

	wrapperlog "github.com/Lucino772/envelop/internal/wrapper/log"
)

type EventsLoggingHandler struct {
	wrapper Wrapper
}

func NewEventsHandler(w Wrapper) *EventsLoggingHandler {
	return &EventsLoggingHandler{wrapper: w}
}

func (handler *EventsLoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (handler *EventsLoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return handler
}

func (handler *EventsLoggingHandler) WithGroup(name string) slog.Handler {
	return handler
}

func (handler *EventsLoggingHandler) Handle(ctx context.Context, record slog.Record) error {
	level := wrapperlog.LevelAttributeReplacer([]string{}, slog.Any(slog.LevelKey, record.Level))
	event := LogEvent{
		Time:    record.Time,
		Message: record.Message,
		Level:   level.Value.String(),
		Data:    make(map[string]any),
	}
	record.Attrs(func(a slog.Attr) bool {
		event.Data[a.Key] = a.Value.Any()
		return true
	})
	handler.wrapper.EmitEvent(event)
	return nil
}
