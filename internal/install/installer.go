package install

import (
	"context"
	"log"
	"os"
	"sync"
)

type Getter interface {
	Get(context.Context, string) error
}

type Decompressor interface {
	Decompress(context.Context, string, string) error
}

type InstallProcessor interface {
	WithInstallDir(string) InstallProcessor
	GetSize() uint32
	IterTasks(yield func(*InstallTask) bool)
}

type Installer struct {
	processors []InstallProcessor
}

func NewInstaller(processors []InstallProcessor) *Installer {
	return &Installer{
		processors: processors,
	}
}

func (installer *Installer) Install(ctx context.Context, installDir string) {
	taskQueue := make(chan *InstallTask, 20)
	go func() {
		defer close(taskQueue)
		for _, processor := range installer.processors {
			processor.WithInstallDir(installDir).IterTasks(func(task *InstallTask) bool {
				select {
				case taskQueue <- task:
					return true
				case <-ctx.Done():
					return false
				}
			})
		}
	}()

	var wg sync.WaitGroup
	for task := range taskQueue {
		wg.Add(1)
		go func(ctx context.Context, task *InstallTask, wg *sync.WaitGroup) {
			defer wg.Done()
			if err := task.Run(ctx); err != nil {
				log.Printf("An error occured %v\n", err)
			}
		}(ctx, task, &wg)
	}
	wg.Wait()
}

type InstallTask struct {
	Path         string
	Getter       Getter
	Decompressor Decompressor
}

func (task *InstallTask) Run(ctx context.Context) error {
	dstPath := task.Path
	if task.Decompressor != nil {
		tmpFile, err := os.CreateTemp("", "")
		if err != nil {
			if tmpFile != nil {
				tmpFile.Close()
			}
			return err
		}
		dstPath = tmpFile.Name()
		tmpFile.Close()
	}
	if err := task.Getter.Get(ctx, dstPath); err != nil {
		return err
	}
	if task.Decompressor != nil {
		if err := task.Decompressor.Decompress(ctx, dstPath, task.Path); err != nil {
			return err
		}
	}
	return nil
}
