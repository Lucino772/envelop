package configs

import (
	"io/fs"

	"github.com/Lucino772/envelop/internal/wrapper"
	"github.com/magiconair/properties"
	"github.com/mitchellh/mapstructure"
)

type propertiesParser struct {
	Filename string `mapstructure:"filename,omitempty"`
}

func newPropertiesParser(opts map[string]any) wrapper.ConfigParser {
	var parser propertiesParser
	if err := mapstructure.Decode(opts, &parser); err != nil {
		return nil
	}
	return &parser
}

func (parser *propertiesParser) Parse(config map[string]any, wp wrapper.Wrapper) error {
	data, err := fs.ReadFile(wp.Files(), parser.Filename)
	if err != nil {
		return err
	}
	props, err := properties.LoadString(string(data))
	if err != nil {
		return err
	}
	for key, val := range props.Map() {
		config[key] = val
	}
	return nil
}
