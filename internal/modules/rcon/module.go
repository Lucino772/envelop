package rcon

import (
	"github.com/Lucino772/envelop/internal/wrapper"
	"github.com/mitchellh/mapstructure"
)

func Initialize(opts map[string]any, registry *wrapper.Registry) {
	var config struct {
		PasswordConfigKey string `mapstructure:"password_key"`
		PortConfigKey     string `mapstructure:"port_key"`
		EnabledConfigKey  string `mapstructure:"enabled_key,omitempty"`
	}
	if err := mapstructure.Decode(opts, &config); err != nil {
		return
	}
	registry.Services = append(
		registry.Services,
		newRconService(
			config.PasswordConfigKey,
			config.PortConfigKey,
			config.EnabledConfigKey,
		),
	)
}
