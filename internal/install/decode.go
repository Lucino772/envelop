package install

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

var (
	ErrMissingSourceType = errors.New("missing type attribute in source")
	ErrUnknownSourceType = errors.New("unknown source type")
)

func decodeSource(data map[string]any) (Source, error) {
	sType, ok := data["type"]
	if !ok {
		return nil, ErrMissingSourceType
	}

	switch sType.(string) {
	case "url":
		return decode(data, &UrlSource{})
	case "http":
		return decode(data, &HttpSource{})
	case "base64":
		return decode(data, &Base64Source{})
	default:
		return nil, ErrUnknownSourceType
	}
}

func decode(data map[string]any, target Source) (Source, error) {
	var decoderMD mapstructure.Metadata
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: &decoderMD,
		TagName:  "mapstructure",
		Result:   target,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		return nil, err
	}
	return target, nil
}
