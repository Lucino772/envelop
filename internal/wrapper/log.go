package wrapper

import (
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

func NewLoggingHandler(typ string, opts map[string]any) slog.Handler {
	switch typ {
	case "default":
		return NewDefaultLoggingHandler(opts)
	default:
		return nil
	}
}

func levelAttributeReplacer(groups []string, a slog.Attr) slog.Attr {
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
