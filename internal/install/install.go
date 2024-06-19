package install

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/sync/errgroup"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Getter interface {
	Get(context.Context, string) error
}

type Decompressor interface {
	Decompress(context.Context, string, string) error
}

type SourceProcessor interface {
	IterTasks(src Source, yield func(*SourceProcessorTask) bool)
	GetExportVars(src Source) any
}

type SourceProcessorTask struct {
	Path         string
	Getter       Getter
	Decompressor Decompressor
}

type Installer struct {
	sourceProcessors map[string]SourceProcessor
}

func NewInstaller() *Installer {
	return &Installer{
		sourceProcessors: map[string]SourceProcessor{
			"http":  &HttpProcessor{},
			"https": &HttpProcessor{},
			"data":  &DataProcessor{},
		},
	}
}

func (i *Installer) RegisterSourceProcessor(name string, processor SourceProcessor) {
	i.sourceProcessors[name] = processor
}

func (i *Installer) Install(ctx context.Context, m *Manifest, directory string) error {
	m = m.WithInstallDir(directory)
	exports := make(map[string]any, 0)

	errg, newCtx := errgroup.WithContext(ctx)
	errg.SetLimit(10)
	for _, source := range m.Sources {
		if ctx.Err() != nil {
			break
		}

		processor, ok := i.sourceProcessors[source.Url.Scheme]
		if !ok {
			return errors.New("unknown source schema")
		}
		processor.IterTasks(source, func(task *SourceProcessorTask) bool {
			errg.Go(func() error {
				return task.run(newCtx)
			})
			return newCtx.Err() == nil
		})

		for key, val := range parseExports(source.Exports, processor.GetExportVars(source)) {
			exports[key] = val
		}
	}
	if err := errg.Wait(); err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(directory, "envelop.yaml"))
	if err != nil {
		if file != nil {
			file.Close()
		}
		return err
	}
	defer file.Close()

	content, err := gameConfigs.ReadFile(filepath.Join("data/configs", m.Config))
	if err != nil {
		return err
	}
	tmpl, err := template.New(m.Config).Parse(string(content))
	if err != nil {
		return err
	}
	return tmpl.Execute(file, exports)
}

func (task *SourceProcessorTask) run(ctx context.Context) error {
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
