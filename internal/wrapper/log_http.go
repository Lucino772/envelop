package wrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

type HttpLoggingHandler struct {
	Url   string     `mapstructure:"url,omitempty"`
	Level slog.Level `mapstructure:"level,omitempty"`
}

type httpLogData struct {
	Time    time.Time      `json:"timestamp"`
	Message string         `json:"message"`
	Level   string         `json:"level"`
	Data    map[string]any `json:"data"`
}

func NewHttpLoggingHandler(opts map[string]any) slog.Handler {
	var handler HttpLoggingHandler
	if err := mapstructure.Decode(opts, &handler); err != nil {
		return nil
	}
	return &handler
}

func (handler *HttpLoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// TODO: Make this configurable
	if level == LevelProcess {
		return false
	}
	return level >= handler.Level
}

func (handler *HttpLoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return handler
}

func (handler *HttpLoggingHandler) WithGroup(name string) slog.Handler {
	return handler
}

func (handler *HttpLoggingHandler) Handle(parent context.Context, record slog.Record) error {
	level := levelAttributeReplacer([]string{}, slog.Any(slog.LevelKey, record.Level))
	event := httpLogData{
		Time:    record.Time,
		Message: record.Message,
		Level:   level.Value.String(),
		Data:    make(map[string]any),
	}
	record.Attrs(func(a slog.Attr) bool {
		event.Data[a.Key] = a.Value.Any()
		return true
	})

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// TODO: What about security/authentication
	ctx, cancel := context.WithTimeout(parent, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", handler.Url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: Do we except a response ? If so, what's the shape ?
	return nil
}
