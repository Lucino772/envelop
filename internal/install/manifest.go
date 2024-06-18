package install

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"text/template"

	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/sync/semaphore"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var ErrManifestNotExists = errors.New("manifest does not exists")

//go:embed manifest-spec.json
var manifestSchema string

//go:embed data/manifest.json
var manifestData []byte

//go:embed data/configs/*
var gameConfigs embed.FS

type Getter interface {
	Get(context.Context, string) error
}

type Decompressor interface {
	Decompress(context.Context, string, string) error
}

type InstallProcessor interface {
	WithInstallDir(string) InstallProcessor
	ParseExports() map[string]interface{}
	GetSize() uint32
	IterTasks(yield func(*InstallTask) bool)
}

type InstallTask struct {
	Path         string
	Getter       Getter
	Decompressor Decompressor
}

type Manifest struct {
	Sources []InstallProcessor `json:"sources,omitempty"`
	Config  string             `json:"config,omitempty"`
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

func (m *Manifest) Install(ctx context.Context, installDir string) error {
	processors := make([]InstallProcessor, 0)
	for _, processor := range m.Sources {
		processors = append(processors, processor.WithInstallDir(installDir))
	}

	var wg sync.WaitGroup
	var sem = semaphore.NewWeighted(10)
	for _, processor := range processors {
		processor.IterTasks(func(task *InstallTask) bool {
			if err := sem.Acquire(ctx, 1); err != nil {
				return false
			}
			defer sem.Release(1)

			wg.Add(1)
			go func(ctx context.Context, task *InstallTask, wg *sync.WaitGroup) {
				defer wg.Done()
				if err := task.run(ctx); err != nil {
					log.Printf("An error occured %v\n", err)
				}
			}(ctx, task, &wg)
			return true
		})
		if ctx.Err() != nil {
			break
		}
	}
	wg.Wait()

	file, err := os.Create(filepath.Join(installDir, "envelop.yaml"))
	if err != nil {
		if file != nil {
			file.Close()
		}
		return err
	}
	defer file.Close()

	exports := make(map[string]any, 0)
	for _, processor := range processors {
		for key, val := range processor.ParseExports() {
			exports[key] = val
		}
	}
	content, err := gameConfigs.ReadFile(filepath.Join("data/configs", m.Config))
	if err != nil {
		return err
	}
	tmpl, err := template.New(m.Config).Parse(string(content))
	if err != nil {
		return err
	}
	return tmpl.Execute(file, exports)
}

func (task *InstallTask) run(ctx context.Context) error {
	dstPath := task.Path
	if task.Decompressor != nil {
		tmpFile, err := os.CreateTemp("", "")
		if err != nil {
			if tmpFile != nil {
				tmpFile.Close()
			}
			return err
		}
		dstPath = tmpFile.Name()
		tmpFile.Close()
	}
	if err := task.Getter.Get(ctx, dstPath); err != nil {
		return err
	}
	if task.Decompressor != nil {
		if err := task.Decompressor.Decompress(ctx, dstPath, task.Path); err != nil {
			return err
		}
	}
	return nil
}

func manifestDecodeHook(typ reflect.Type, target reflect.Type, data any) (any, error) {
	if typ.Kind() == reflect.Map && target == reflect.TypeOf((*InstallProcessor)(nil)).Elem() {
		decoders := map[string]func(map[string]interface{}) (InstallProcessor, error){
			"files":   decodeFilesSource,
			"archive": decodeArchiveSource,
			"content": decodeContentSource,
		}
		sourceData := data.(map[string]any)
		sourceType := sourceData["type"].(string)
		decoder, ok := decoders[sourceType]
		if !ok {
			return nil, errors.New("invalid decoder")
		}
		return decoder(sourceData)
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
