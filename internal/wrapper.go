package wrapper

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-cmd/cmd"
)

type WrapperStopper func(*Wrapper) error

type Wrapper struct {
	Dir     string
	Env     []string
	Timeout time.Duration

	cmd             *cmd.Cmd
	stdinReader     io.Reader
	stdinWriter     io.WriteCloser
	gracefulStopper WrapperStopper
}

func NewGracefulSignalWrapper(stopSignal os.Signal, program string, args ...string) (*Wrapper, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	return &Wrapper{
		Timeout: 30 * time.Second,
		cmd: cmd.NewCmdOptions(cmd.Options{
			Buffered:  false,
			Streaming: true,
		}, program, args...),
		stdinReader:     stdinReader,
		stdinWriter:     stdinWriter,
		gracefulStopper: withGracefulSignal(stopSignal),
	}, nil
}

func NewGracefulCommandWrapper(stopCommand string, program string, args ...string) (*Wrapper, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	return &Wrapper{
		Timeout: 30 * time.Second,
		cmd: cmd.NewCmdOptions(cmd.Options{
			Buffered:  false,
			Streaming: true,
		}, program, args...),
		stdinReader:     stdinReader,
		stdinWriter:     stdinWriter,
		gracefulStopper: WithGracefulCommand(stopCommand),
	}, nil
}

func (wrapper *Wrapper) Write(value string) (int, error) {
	return wrapper.stdinWriter.Write([]byte(fmt.Sprintf("%s\n", value)))
}

func (wrapper *Wrapper) Run() {
	defer wrapper.stdinWriter.Close()
	wrapper.cmd.Dir = wrapper.Dir
	wrapper.cmd.Env = wrapper.Env

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	statusChan := wrapper.cmd.StartWithStdin(wrapper.stdinReader)
	go wrapper.produceLogs()

	select {
	case <-signalChan:
		wrapper.gracefulStop(statusChan)
		wrapper.cmd.Stop()
	case <-statusChan:
		signal.Stop(signalChan)
	}
}

func withGracefulSignal(signal os.Signal) WrapperStopper {
	return func(wrapper *Wrapper) error {
		process, err := os.FindProcess(wrapper.cmd.Status().PID)
		if err != nil {
			return err
		}
		process.Signal(signal)
		return nil
	}
}

func WithGracefulCommand(command string) WrapperStopper {
	return func(wrapper *Wrapper) error {
		_, err := wrapper.Write(command)
		return err
	}
}

func (wrapper *Wrapper) gracefulStop(statusChan <-chan cmd.Status) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wrapper.Timeout)
	defer cancel()

	err := wrapper.gracefulStopper(wrapper)
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

func (wrapper *Wrapper) produceLogs() {
	// TODO: Replace log with pubsub model
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for value := range wrapper.cmd.Stdout {
			log.Println(value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range wrapper.cmd.Stderr {
			log.Println(value)
		}
	}()
	wg.Wait()
}
