package install

import (
	"embed"
	"encoding/json"
	"errors"
	"net/url"
	"path/filepath"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
)

var ErrManifestNotExists = errors.New("manifest does not exists")

//go:embed data/manifest-spec.json
var manifestSchema string

//go:embed data/manifest.json
var manifestData []byte

//go:embed data/configs/*
var gameConfigs embed.FS

type Source struct {
	Url         url.URL        `json:"url,omitempty"`
	Destination string         `json:"destination,omitempty"`
	Exports     map[string]any `json:"exports,omitempty"`
}

type Manifest struct {
	Sources []Source `json:"sources,omitempty"`
	Config  string   `json:"config,omitempty"`
}

func GetManifest(id string) (*Manifest, error) {
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
		TagName:    "json",
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

func (m *Manifest) WithInstallDir(dir string) *Manifest {
	sources := make([]Source, 0)
	for _, src := range m.Sources {
		dest, err := filepath.Abs(filepath.Join(dir, src.Destination))
		if err != nil {
			dest = filepath.Join(dir, src.Destination)
		}

		sources = append(sources, Source{
			Url:         src.Url,
			Destination: dest,
			Exports:     src.Exports,
		})
	}
	return &Manifest{
		Sources: sources,
		Config:  m.Config,
	}
}

func manifestDecodeHook(typ reflect.Type, target reflect.Type, data any) (any, error) {
	if typ.Kind() == reflect.String && target == reflect.TypeOf((*url.URL)(nil)).Elem() {
		return url.Parse(data.(string))
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
