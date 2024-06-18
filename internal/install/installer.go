package install

import (
	"context"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Getter interface {
	Get(context.Context, string) error
}

type Decompressor interface {
	Decompress(context.Context, string, string) error
}

type InstallProcessor interface {
	WithInstallDir(string) InstallProcessor
	ParseExports() map[string]interface{}
	GetSize() uint32
	IterTasks(yield func(*InstallTask) bool)
}

type Installer struct {
	processors []InstallProcessor
	config     string
}

func NewInstaller(processors []InstallProcessor, config string) *Installer {
	return &Installer{
		processors: processors,
		config:     config,
	}
}

func (installer *Installer) Install(ctx context.Context, installDir string) {
	taskQueue := make(chan *InstallTask, 20)
	processors := make([]InstallProcessor, 0)
	for _, processor := range installer.processors {
		processors = append(processors, processor.WithInstallDir(installDir))
	}

	go func() {
		defer close(taskQueue)
		for _, processor := range processors {
			processor.IterTasks(func(task *InstallTask) bool {
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

	if err := installer.createConfig(processors, installDir); err != nil {
		log.Printf("Failed to create config")
	}
}

func (installer *Installer) createConfig(processors []InstallProcessor, installDir string) error {
	tmpl, err := template.New("config").Parse(installer.config)
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(installDir, "envelop.yaml"))
	if err != nil {
		if file != nil {
			file.Close()
		}
		return err
	}
	defer file.Close()

	exports := make(map[string]any, 0)
	for _, processor := range processors {
		for key, val := range processor.ParseExports() {
			exports[key] = val
		}
	}
	return tmpl.Execute(file, exports)
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

func parseExports(exports map[string]any, data any) map[string]any {
	exp := make(map[string]any, 0)
	for key, value := range exports {
		formattedKey := cases.Title(language.English, cases.Compact).String(strings.ToLower(key))
		if stringVal, ok := value.(string); ok {
			templ, err := template.New(key).Parse(stringVal)
			if err != nil {
				continue
			}
			var buf strings.Builder
			if err := templ.Execute(&buf, data); err != nil {
				continue
			}
			exp[formattedKey] = buf.String()
		} else {
			exp[formattedKey] = value
		}
	}
	return exp
}
