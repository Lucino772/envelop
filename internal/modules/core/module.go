package core

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any) wrapper.Module {
	return func(options *wrapper.Options) {
		options.Services = append(
			options.Services,
			NewCoreSystemService(),
			NewCoreProcessService(),
			NewCorePlayersService(),
		)
	}
}
