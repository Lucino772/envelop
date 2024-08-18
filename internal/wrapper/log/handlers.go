package wrapperlog

import (
	"io"
	"log/slog"
	"os"

	"github.com/mitchellh/mapstructure"
)

func NewHandler(name string, opts map[string]any) slog.Handler {
	switch name {
	case "default":
		var options struct {
			Format string     `mapstructure:"format"`
			Level  slog.Level `mapstructure:"level"`
		}
		if err := mapstructure.Decode(opts, &options); err != nil {
			return nil
		}
		if options.Format == "" {
			options.Format = "text"
		}

		if options.Format == "" || options.Format == "text" {
			return NewTextHandler(os.Stdout, options.Level)
		} else if options.Format == "json" {
			return NewJSONHandler(os.Stdout, options.Level)
		} else {
			return nil
		}
	default:
		return nil
	}
}

func NewJSONHandler(w io.Writer, level slog.Level) slog.Handler {
	return slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level:       level,
		ReplaceAttr: LevelAttributeReplacer,
	})
}

func NewTextHandler(w io.Writer, level slog.Level) slog.Handler {
	return slog.NewTextHandler(w, &slog.HandlerOptions{
		Level:       level,
		ReplaceAttr: LevelAttributeReplacer,
	})
}
