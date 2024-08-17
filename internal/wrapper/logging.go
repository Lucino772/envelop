package wrapper

import (
	"context"
	"fmt"
	"log/slog"
)

var (
	LevelDebug   = slog.LevelDebug
	LevelInfo    = slog.LevelInfo
	LevelWarn    = slog.LevelWarn
	LevelError   = slog.LevelError
	LevelProcess = slog.Level(12)
)

func attributeReplace(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)

		str := func(base string, val slog.Level) slog.Value {
			if val == 0 {
				return slog.StringValue(base)
			}
			return slog.StringValue(fmt.Sprintf("%s%+d", base, val))
		}

		switch {
		case level < LevelInfo:
			a.Value = str("DEBUG", level-LevelDebug)
		case level < LevelWarn:
			a.Value = str("INFO", level-LevelInfo)
		case level < LevelError:
			a.Value = str("WARN", level-LevelWarn)
		case level < LevelProcess:
			a.Value = str("ERROR", level-LevelError)
		default:
			a.Value = str("PROCESS", level-LevelProcess)
		}
	}
	return a
}

type LoggingHandler struct {
	wrapper Wrapper
}

func NewLoggingHandler(wrapper Wrapper) *LoggingHandler {
	return &LoggingHandler{
		wrapper: wrapper,
	}
}

func (handler *LoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (handler *LoggingHandler) Handle(ctx context.Context, record slog.Record) error {
	level := attributeReplace([]string{}, slog.Any(slog.LevelKey, record.Level))
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

func (handler *LoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return handler
}

func (handler *LoggingHandler) WithGroup(name string) slog.Handler {
	return handler
}
