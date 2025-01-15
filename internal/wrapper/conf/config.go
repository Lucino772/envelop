package wrapperconf

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/Lucino772/envelop/internal/modules"
	"github.com/Lucino772/envelop/internal/utils/logutils"
	"github.com/Lucino772/envelop/internal/wrapper"
	"github.com/google/shlex"
	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

var ErrInvalidWrapperConfig = errors.New("invalid wrapper config")

//go:embed envelop-spec.json
var configSchema string

type configData struct {
	Process struct {
		Command     string            `yaml:"command,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
		Graceful    struct {
			Type    string                 `yaml:"type,omitempty"`
			Timeout int                    `yaml:"timeout,omitempty"`
			Options map[string]interface{} `yaml:"options,omitempty"`
		} `yaml:"graceful,omitempty"`
	} `yaml:"process,omitempty"`
	Hooks []struct {
		Type    string         `yaml:"type,omitempty"`
		Options map[string]any `yaml:"options,omitempty"`
	} `yaml:"hooks,omitempty"`
	Logging []struct {
		Type    string         `yaml:"type,omitempty"`
		Options map[string]any `yaml:"options,omitempty"`
	} `yaml:"logging,omitempty"`
	Modules []struct {
		Name    string                 `yaml:"uses,omitempty"`
		Options map[string]interface{} `yaml:"with,omitempty"`
	} `yaml:"modules,omitempty"`
}

func LoadFile(path string) (*wrapper.Options, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Load(data)
}

func Load(source []byte) (*wrapper.Options, error) {
	var dict map[string]interface{}
	yaml.Unmarshal(source, &dict)

	if err := validateConfig(dict); err != nil {
		return nil, err
	}
	data, err := decodeConfig(dict)
	if err != nil {
		return nil, err
	}

	var options wrapper.Options

	command, err := shlex.Split(data.Process.Command)
	if err != nil {
		return nil, err
	}
	options.Program = command[0]
	options.Args = command[1:]

	options.Graceful.Timeout = time.Duration(data.Process.Graceful.Timeout) * time.Second
	options.Graceful.Stopper = wrapper.NewGracefulStopper(data.Process.Graceful.Type, data.Process.Graceful.Options)

	for _, hook := range data.Hooks {
		h := wrapper.NewHook(hook.Type, hook.Options)
		if h != nil {
			// TODO: Add name from hook
			options.Tasks = append(
				options.Tasks,
				wrapper.NewNamedTask(
					"hook",
					func(ctx context.Context, wp wrapper.Wrapper) error {
						sub := wp.SubscribeEvents()
						defer sub.Close()

						for event := range sub.Receive() {
							data, err := json.Marshal(event)
							if err == nil {
								// TODO: Handle error, log maybe ?
								_ = h.Execute(ctx, data)
							}
						}
						return nil
					},
				),
			)
		}
	}

	handlers := make([]slog.Handler, 0)
	for _, logconf := range data.Logging {
		if handler := wrapper.NewLoggingHandler(logconf.Type, logconf.Options); handler != nil {
			handlers = append(handlers, handler)
		}
	}
	options.Logger = slog.New(logutils.NewMultiHandler(handlers...))

	for _, mod := range data.Modules {
		if module := modules.NewModule(mod.Name, mod.Options); module != nil {
			module(&options)
		}
	}
	return &options, nil
}

func validateConfig(config map[string]interface{}) error {
	schemaLoader := gojsonschema.NewStringLoader(configSchema)
	dataLoader := gojsonschema.NewGoLoader(config)

	res, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		return err
	}

	if !res.Valid() {
		return ErrInvalidWrapperConfig
	}
	return nil
}

func decodeConfig(config map[string]any) (*configData, error) {
	var conf configData
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
	if err := decoder.Decode(config); err != nil {
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
