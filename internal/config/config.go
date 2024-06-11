package config

import (
	"errors"
	"fmt"
	"slices"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Process Process        `yaml:"process,omitempty"`
	Modules []ModuleConfig `yaml:"modules,omitempty"`
}

type Process struct {
	Command     string            `yaml:"command,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Graceful    struct {
		Type    string                 `yaml:"type,omitempty"`
		Timeout int                    `yaml:"timeout,omitempty"`
		Options map[string]interface{} `yaml:"options,omitempty"`
	} `yaml:"graceful,omitempty"`
}

type ModuleConfig struct {
	Uses string                 `yaml:"uses,omitempty"`
	With map[string]interface{} `yaml:"with,omitempty"`
}

func LoadConfig(source []byte) (*Config, error) {
	var dict map[string]interface{}
	yaml.Unmarshal(source, &dict)

	if err := Validate(dict); err != nil {
		return nil, err
	}

	var conf Config
	var decoderMD mapstructure.Metadata
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: &decoderMD,
		TagName:  "yaml",
		Result:   &conf,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(dict); err != nil {
		return nil, err
	}

	if !slices.Contains([]string{"cmd", "signal"}, conf.Process.Graceful.Type) {
		return nil, fmt.Errorf("%s not recognized as a graceful type", conf.Process.Graceful.Type)
	}
	switch conf.Process.Graceful.Type {
	case "cmd":
		if _, ok := conf.Process.Graceful.Options["cmd"]; !ok {
			return nil, errors.New("missing `cmd` in graceful options")
		}
	case "signal":
		if _, ok := conf.Process.Graceful.Options["signal"]; !ok {
			return nil, errors.New("missing `signal` in graceful options")
		}
	}

	return &conf, nil
}
