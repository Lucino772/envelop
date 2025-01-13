package install

import (
	_ "embed"
	"log"
	"path/filepath"
	"slices"
	"text/template"

	"context"
	"errors"
	"os"

	"github.com/Lucino772/envelop/pkg/download"
	"github.com/alitto/pond/v2"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

var ErrManifestNotExists = errors.New("manifest does not exists")

type Installer struct{}

func NewInstaller() (*Installer, error) {
	return &Installer{}, nil
}

func (i *Installer) GetManifest(ctx context.Context, uri string) (*Manifest, error) {
	outFile, err := os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}
	outFile.Close()
	defer os.Remove(outFile.Name())

	dl := download.NewDownloader(uri, outFile.Name())
	if err := dl.Download(ctx); err != nil {
		return nil, err
	}

	manifestData, err := os.ReadFile(outFile.Name())
	if err != nil {
		return nil, err
	}

	var manifestMap map[string]any
	if err := yaml.Unmarshal(manifestData, &manifestMap); err != nil {
		return nil, err
	}
	if err := validateManifest(manifestMap); err != nil {
		return nil, err
	}

	var manifest Manifest
	if _, err := decode(manifestMap, &manifest); err != nil {
		return nil, err
	}
	return &manifest, nil
}

func (i *Installer) Install(ctx context.Context, m *Manifest, config DownloadConfig) error {
	depots := make([]Depot, 0)
	exports := make(map[string]any, 0)

	var dlOptions []DownloaderOptFunc
	dlOptions = append(dlOptions, WithDownloadConfig(config))
	for _, depot := range m.Depots {
		if !slices.Contains(depot.Config.Os, "*") && !slices.Contains(depot.Config.Os, config.TargetOs) {
			continue
		}
		if !slices.Contains(depot.Config.Arch, "*") && !slices.Contains(depot.Config.Arch, config.TargetArch) {
			continue
		}
		depot.Path = filepath.ToSlash(filepath.Join(config.InstallDir, depot.Path))

		vars := struct {
			Path       string
			TargetOs   string
			TargetArch string
		}{
			Path:       depot.Path,
			TargetOs:   config.TargetOs,
			TargetArch: config.TargetArch,
		}
		for key, val := range parseExports(depot.Exports, vars) {
			exports[key] = val
		}
		log.Println(exports)

		dlOptions = append(dlOptions, depot.Manifest.GetDownloaderOptions()...)
		depots = append(depots, depot)
	}

	downloader := NewDownloader(dlOptions...)
	if err := downloader.Initialize(); err != nil {
		return err
	}

	metadatas := make([]Metadata, 0)
	for _, depot := range depots {
		metadata, err := depot.Manifest.GetMetadata(ctx, downloader, depot.Path)
		if err != nil {
			return err
		}
		if metadata != nil {
			metadatas = append(metadatas, metadata)
		}
	}

	errg, errCtx := errgroup.WithContext(ctx)
	workerPool := pond.NewPool(12, pond.WithContext(errCtx))
	for _, metadata := range metadatas {
		errg.Go(func() error {
			waiter, err := metadata.Install(errCtx, workerPool, downloader)
			if err != nil {
				return err
			}
			return waiter.Wait()
		})
	}
	err := errg.Wait()
	workerPool.StopAndWait()
	if err != nil {
		return err
	}

	// Create config
	configPath := filepath.Join(config.InstallDir, "envelop.yaml")
	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	tmpl, err := template.New(configPath).Parse(m.Config)
	if err != nil {
		return err
	}
	return tmpl.Execute(configFile, exports)
}
