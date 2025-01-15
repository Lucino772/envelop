package wrapper

import (
	"context"
	"errors"
	"log/slog"
)

type task struct {
	name string
	run  func(context.Context, Wrapper) error
}

func NewNamedTask(name string, run func(context.Context, Wrapper) error) *task {
	return &task{
		name: name,
		run:  run,
	}
}

func (task *task) Name() string {
	return task.name
}

func (task *task) Run(ctx context.Context, wp Wrapper) error {
	return task.run(ctx, wp)
}

func makeRecoverableTask(ctx context.Context, task Task, logger *slog.Logger, wp Wrapper) func() error {
	return func() (err error) {
		taskLogger := logger.With(slog.String("task", task.Name()))
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
			if err != nil {
				taskLogger.LogAttrs(
					ctx,
					LevelError,
					"task stopped with error",
					slog.Any("error", err),
				)
			} else {
				taskLogger.LogAttrs(ctx, LevelInfo, "task done")
			}
		}()

		taskLogger.LogAttrs(ctx, LevelInfo, "task started")
		if r := task.Run(ctx, wp); r != nil && !errors.Is(r, context.Canceled) {
			err = r
		}
		return err
	}
}
