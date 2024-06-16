package install

import (
	_ "embed"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

//go:embed install-manifest.json
var manifestsData []byte

type Manifest struct {
	Sources []InstallProcessor
	Config  string
}

func LoadManifestConfig(manifestId string) (*Manifest, error) {
	var dict map[string]map[string]interface{}
	json.Unmarshal(manifestsData, &dict)

	manifestDict, ok := dict[manifestId]
	if !ok {
		return nil, errors.New("manifest does not exist")
	}
	if err := Validate(manifestDict); err != nil {
		return nil, err
	}

	var manifestConf struct {
		Sources []map[string]interface{} `json:"sources,omitempty"`
		Config  string                   `json:"config,omitempty"`
	}
	if err := decode(manifestDict, &manifestConf); err != nil {
		return nil, err
	}

	manifest := Manifest{
		Sources: make([]InstallProcessor, 0),
		Config:  manifestConf.Config,
	}
	decoders := map[string]func(map[string]interface{}) (InstallProcessor, error){
		"files":   decodeFilesSource,
		"archive": decodeArchiveSource,
		"content": decodeContentSource,
	}
	for _, source := range manifestConf.Sources {
		sourceType := source["type"].(string)
		if decoder, ok := decoders[sourceType]; ok {
			config, err := decoder(source)
			if err != nil {
				return nil, err
			}
			manifest.Sources = append(manifest.Sources, config)
		}
	}

	return &manifest, nil
}

func decode(source map[string]interface{}, target interface{}) error {
	var decoderMD mapstructure.Metadata
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: &decoderMD,
		TagName:  "json",
		Result:   target,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return err
	}
	if err := decoder.Decode(source); err != nil {
		return err
	}
	return nil
}

func decodeFilesSource(source map[string]interface{}) (InstallProcessor, error) {
	var conf FilesProcessor
	if err := decode(source, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func decodeArchiveSource(source map[string]interface{}) (InstallProcessor, error) {
	var conf ArchiveProcessor
	if err := decode(source, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func decodeContentSource(source map[string]interface{}) (InstallProcessor, error) {
	var conf ContentProcessor
	if err := decode(source, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
