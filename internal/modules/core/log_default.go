package core

import (
	"log/slog"
	"os"

	"github.com/mitchellh/mapstructure"
)

func NewDefaultLoggingHandler(opts map[string]any) slog.Handler {
	var config struct {
		Format string     `mapstructure:"format"`
		Level  slog.Level `mapstructure:"level"`
	}
	if err := mapstructure.Decode(opts, &config); err != nil {
		return nil
	}
	if config.Format == "" {
		config.Format = "text"
	}

	handlerOpts := &slog.HandlerOptions{
		Level:       config.Level,
		ReplaceAttr: levelAttributeReplacer,
	}
	switch config.Format {
	case "text":
		return slog.NewTextHandler(os.Stdout, handlerOpts)
	case "json":
		return slog.NewJSONHandler(os.Stdout, handlerOpts)
	default:
		return nil
	}
}
