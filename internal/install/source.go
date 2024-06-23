package install

import (
	"errors"
	"strings"
	"text/template"

	"github.com/Lucino772/envelop/pkg/download"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	ErrMissingSourceType = errors.New("missing type attribute in source")
	ErrUnknownSourceType = errors.New("unknown source type")
)

type Source interface {
	GetExports() map[string]any
	WithInstallDir(dir string) Source
	IterTasks(yield func(*download.Downloader) bool)
}

func decodeSource(data map[string]any) (Source, error) {
	decoders := map[string]func(map[string]any) (Source, error){
		"url": decodeUrlSource,
	}

	sType, ok := data["type"]
	if !ok {
		return nil, ErrMissingSourceType
	}
	decoder, ok := decoders[sType.(string)]
	if !ok {
		return nil, ErrUnknownSourceType
	}
	return decoder(data)
}

func decodeUrlSource(data map[string]any) (Source, error) {
	var source UrlSource
	if err := decode(data, &source); err != nil {
		return nil, err
	}
	return &source, nil
}

func decode(data map[string]any, target any) error {
	var decoderMD mapstructure.Metadata
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: &decoderMD,
		TagName:  "mapstructure",
		Result:   target,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return err
	}
	if err := decoder.Decode(data); err != nil {
		return err
	}
	return nil
}

func parseExports(exports map[string]any, data any) map[string]any {
	exp := make(map[string]any, 0)
	for key, value := range exports {
		formattedKey := cases.Title(language.English, cases.Compact).String(strings.ToLower(key))
		if stringVal, ok := value.(string); ok {
			templ, err := template.New(key).Parse(stringVal)
			if err != nil {
				continue
			}
			var buf strings.Builder
			if err := templ.Execute(&buf, data); err != nil {
				continue
			}
			exp[formattedKey] = buf.String()
		} else {
			exp[formattedKey] = value
		}
	}
	return exp
}
