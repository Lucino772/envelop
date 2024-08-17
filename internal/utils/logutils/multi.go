package logutils

import (
	"context"
	"errors"
	"log/slog"
	"slices"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	var result error
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, record.Level) {
			// TODO: Run handlers in separate goroutines
			if err := handler.Handle(ctx, record.Clone()); err != nil {
				result = errors.Join(err)
			}
		}
	}
	return result
}

func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := &MultiHandler{handlers: make([]slog.Handler, 0)}
	for _, handler := range h.handlers {
		newHandler.handlers = append(
			newHandler.handlers,
			handler.WithAttrs(slices.Clone(attrs)),
		)
	}
	return newHandler
}

func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandler := &MultiHandler{handlers: make([]slog.Handler, 0)}
	for _, handler := range h.handlers {
		newHandler.handlers = append(
			newHandler.handlers,
			handler.WithGroup(name),
		)
	}
	return newHandler
}
