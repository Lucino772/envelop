package install

import (
	"context"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Lucino772/envelop/pkg/download"
	"golang.org/x/sync/errgroup"
)

type Installer struct{}

func NewInstaller() *Installer {
	return &Installer{}
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

		source.IterTasks(func(d *download.Downloader) bool {
			errg.Go(func() error {
				return d.Download(newCtx)
			})
			return newCtx.Err() == nil
		})

		for key, val := range source.GetExports() {
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
