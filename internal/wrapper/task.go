package wrapper

import "context"

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
