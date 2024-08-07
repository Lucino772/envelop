package install

import (
	_ "embed"

	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	"github.com/Lucino772/envelop/pkg/download"
	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/sync/errgroup"
)

var ErrManifestNotExists = errors.New("manifest does not exists")

//go:embed data/manifest-spec.json
var manifestSchema string

type Manifest struct {
	Sources []Source `mapstructure:"sources,omitempty"`
	Config  string   `mapstructure:"config,omitempty"`
}

func (m *Manifest) WithInstallDir(dir string) *Manifest {
	sources := make([]Source, 0)
	for _, src := range m.Sources {
		sources = append(sources, src.WithInstallDir(dir))
	}
	return &Manifest{
		Sources: sources,
		Config:  m.Config,
	}
}

type Installer struct {
	manifestUrl  string
	manifestPath string
}

func NewInstaller() (*Installer, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	return &Installer{
		manifestUrl:  "https://raw.githubusercontent.com/Lucino772/envelop/main/resources/install/manifests.json",
		manifestPath: filepath.Join(userCacheDir, "envelop", "manifests.json"),
	}, nil
}

func (i *Installer) CheckManifestsAvailable() error {
	if _, err := os.Stat(i.manifestPath); err != nil {
		return err
	}
	return nil
}

func (i *Installer) UpdateManifests(ctx context.Context) error {
	cacheDir := filepath.Dir(i.manifestPath)
	if _, err := os.Stat(cacheDir); err != nil {
		os.MkdirAll(cacheDir, os.ModePerm)
	}
	dl := download.NewDownloader(i.manifestUrl, i.manifestPath)
	return dl.Download(ctx)
}

func (i *Installer) GetManifest(id string) (*Manifest, error) {
	if err := i.CheckManifestsAvailable(); err != nil {
		return nil, err
	}
	manifestData, err := os.ReadFile(i.manifestPath)
	if err != nil {
		return nil, err
	}

	var manifests map[string]map[string]any
	if err := json.Unmarshal(manifestData, &manifests); err != nil {
		return nil, err
	}

	data, ok := manifests[id]
	if !ok {
		return nil, ErrManifestNotExists
	}
	if err := validateManifest(data); err != nil {
		return nil, err
	}

	var manifest Manifest
	var decoderMD mapstructure.Metadata
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata:   &decoderMD,
		DecodeHook: manifestDecodeHook,
		TagName:    "mapstructure",
		Result:     &manifest,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		return nil, err
	}
	return &manifest, nil
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

	// TODO: Cache config file to improve performance
	configPath := filepath.Join(directory, "envelop.yaml")
	dl := download.NewDownloader(m.Config, configPath)
	dl.PostDownloadHook = func(dst string) error {
		content, err := os.ReadFile(dst)
		if err != nil {
			return err
		}
		file, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer file.Close()
		tmpl, err := template.New(dst).Parse(string(content))
		if err != nil {
			return err
		}
		return tmpl.Execute(file, exports)
	}
	return dl.Download(ctx)
}

func manifestDecodeHook(typ reflect.Type, target reflect.Type, data any) (any, error) {
	if typ.Kind() == reflect.Map && target == reflect.TypeOf((*Source)(nil)).Elem() {
		return decodeSource(data.(map[string]any))
	}
	return data, nil
}

func validateManifest(config map[string]interface{}) error {
	schemaLoader := gojsonschema.NewStringLoader(manifestSchema)
	dataLoader := gojsonschema.NewGoLoader(config)

	res, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		return err
	}

	if !res.Valid() {
		return errors.New("config is not valid")
	}
	return nil
}
