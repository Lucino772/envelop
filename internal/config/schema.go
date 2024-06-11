package config

import (
	_ "embed"
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed envelop-spec.json
var Schema string

func Validate(config map[string]interface{}) error {
	schemaLoader := gojsonschema.NewStringLoader(Schema)
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
