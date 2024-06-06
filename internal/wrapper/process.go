package wrapper

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-cmd/cmd"
)

type WrapperStopper func(*WrapperProcess) error
type WrapperProcessOpt func(*WrapperProcess)
type WrapperProcessLogCallback func(string)

type WrapperProcess struct {
	args            []string
	dir             string
	env             []string
	logCallback     WrapperProcessLogCallback
	gracefulTimeout time.Duration
	gracefulStopper WrapperStopper

	cmd         *cmd.Cmd
	stdinReader io.Reader
	stdinWriter io.WriteCloser
}

func WithProgramArgs(args ...string) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.args = args
	}
}

func WithDirectory(dir string) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.dir = dir
	}
}

func WithEnv(env []string) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.env = append(wp.env, env...)
	}
}

func WithLogCallback(callback WrapperProcessLogCallback) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.logCallback = callback
	}
}

func WithGracefulTimeout(timeout time.Duration) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.gracefulTimeout = timeout
	}
}

func WithGracefulStopSignal(signal os.Signal) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.gracefulStopper = func(wp *WrapperProcess) error {
			process, err := os.FindProcess(wp.cmd.Status().PID)
			if err != nil {
				return err
			}
			process.Signal(signal)
			return nil
		}
	}
}

func WithGracefulStopCommand(command string) WrapperProcessOpt {
	return func(wp *WrapperProcess) {
		wp.gracefulStopper = func(wp *WrapperProcess) error {
			_, err := wp.Write(command)
			return err
		}
	}
}

func NewWrapperProcess(program string, opts ...WrapperProcessOpt) (*WrapperProcess, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	process := &WrapperProcess{
		args:            make([]string, 0),
		env:             nil,
		gracefulTimeout: 30 * time.Second,
		gracefulStopper: func(wp *WrapperProcess) error { return nil },
		logCallback:     func(s string) {},
		stdinReader:     stdinReader,
		stdinWriter:     stdinWriter,
	}
	for _, opt := range opts {
		opt(process)
	}

	process.cmd = cmd.NewCmdOptions(cmd.Options{
		Buffered:  false,
		Streaming: true,
	}, program, process.args...)
	process.cmd.Dir = process.dir
	process.cmd.Env = process.env
	return process, nil
}

func (wp *WrapperProcess) Write(value string) (int, error) {
	return wp.stdinWriter.Write([]byte(fmt.Sprintf("%s\n", value)))
}

func (wp *WrapperProcess) Run(ctx context.Context) {
	defer wp.stdinWriter.Close()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	statusChan := wp.cmd.StartWithStdin(wp.stdinReader)
	go wp.produceLogs()

	select {
	case <-ctx.Done():
		wp.gracefulStop(statusChan)
		wp.cmd.Stop()
	case <-signalChan:
		wp.gracefulStop(statusChan)
		wp.cmd.Stop()
	case <-statusChan:
		signal.Stop(signalChan)
	}
}

func (wp *WrapperProcess) gracefulStop(statusChan <-chan cmd.Status) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wp.gracefulTimeout)
	defer cancel()

	err := wp.gracefulStopper(wp)
	if err != nil {
		return false, err
	}

	select {
	case <-statusChan:
		return true, nil
	case <-ctx.Done():
		return false, nil
	}
}

func (wp *WrapperProcess) produceLogs() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for value := range wp.cmd.Stdout {
			wp.logCallback(value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range wp.cmd.Stderr {
			wp.logCallback(value)
		}
	}()
	wg.Wait()
}
