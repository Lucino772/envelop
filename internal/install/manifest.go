package install

import (
	"context"
	_ "embed"
	"errors"
	"reflect"

	"github.com/alitto/pond/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
)

//go:embed data/manifest-spec.json
var manifestSchema string

var (
	ErrMissingSourceType        = errors.New("missing type attribute in source")
	ErrUnknownSourceType        = errors.New("unknown source type")
	ErrMissingDepotManifestType = errors.New("missing type attribute in depot manifest")
	ErrUnknownDepotManifestType = errors.New("unknown depot manifest type")
)

type Manifest struct {
	Name   string  `mapstructure:"name,omitempty"`
	Config string  `mapstructure:"config,omitempty"`
	Depots []Depot `mapstructure:"depots,omitempty"`
}

type Depot struct {
	Name     string         `mapstructure:"name,omitempty"`
	Path     string         `mapstructure:"path,omitempty"`
	Config   DepotConfig    `mapstructure:"config,omitempty"`
	Exports  map[string]any `mapstructure:"exports,omitempty"`
	Manifest DepotManifest  `mapstructure:"manifest,omitempty"`
}

type DepotConfig struct {
	Os   []string `mapstructure:"os,omitempty"`
	Arch []string `mapstructure:"arch,omitempty"`
	Tags []string `mapstructure:"tags,omitempty"`
}

type DepotManifest interface {
	GetDownloaderOptions() []DownloaderOptFunc
	GetMetadata(context.Context, *Downloader, string) (Metadata, error)
}

type Metadata interface {
	Install(context.Context, pond.Pool, *Downloader) (Waiter, error)
}

type Source interface {
	GetDownloaderOptions() []DownloaderOptFunc
	Download(context.Context, *Downloader, string) error
}

func decode(input map[string]any, target any) (any, error) {
	var decoderMD mapstructure.Metadata
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata:   &decoderMD,
		DecodeHook: manifestDecodeHook,
		TagName:    "mapstructure",
		Result:     target,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(input); err != nil {
		return nil, err
	}
	return target, nil
}

func manifestDecodeHook(typ reflect.Type, target reflect.Type, data any) (any, error) {
	if typ.Kind() != reflect.Map {
		return data, nil
	}

	// Decode source object
	if target == reflect.TypeOf((*Source)(nil)).Elem() {
		_data := data.(map[string]any)
		sType, ok := _data["type"]
		if !ok {
			return nil, ErrMissingSourceType
		}
		delete(_data, "type")
		switch sType.(string) {
		case "http":
			return decode(_data, &HttpSource{})
		case "base64":
			return decode(_data, &Base64Source{})
		default:
			return nil, ErrUnknownSourceType
		}
	}

	// Decode depot manifest object
	if target == reflect.TypeOf((*DepotManifest)(nil)).Elem() {
		_data := data.(map[string]any)
		sType, ok := _data["type"]
		if !ok {
			return nil, ErrMissingDepotManifestType
		}
		delete(_data, "type")
		switch sType.(string) {
		case "files":
			return decode(_data, &FilesDepotManifest{})
		case "steam":
			return decode(_data, &SteamDepotManifest{})
		default:
			return nil, ErrUnknownDepotManifestType
		}
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
