package wrapper

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-cmd/cmd"
)

func (wp *Wrapper) runProcess(ctx context.Context) {
	defer wp.stdinWriter.Close()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
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

func (wp *Wrapper) gracefulStop(statusChan <-chan cmd.Status) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wp.options.gracefulTimeout)
	defer cancel()

	err := wp.options.gracefulStopper(wp)
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

func (wp *Wrapper) produceLogs() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for value := range wp.cmd.Stdout {
			wp.logsProducer.Publish(value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range wp.cmd.Stderr {
			wp.logsProducer.Publish(value)
		}
	}()
	wg.Wait()
}
