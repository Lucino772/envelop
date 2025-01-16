package core

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any, registry *wrapper.Registry) {
	registry.Services = append(
		registry.Services,
		NewCoreSystemService(),
		NewCoreProcessService(),
		NewCorePlayersService(),
	)
}
