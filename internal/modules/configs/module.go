package configs

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any, registry *wrapper.Registry) {
	registry.ConfigParser["properties"] = newPropertiesParser
}
