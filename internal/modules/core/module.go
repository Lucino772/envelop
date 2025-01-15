package core

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any, wrapperOpts *wrapper.Options) {
	wrapperOpts.Services = append(
		wrapperOpts.Services,
		NewCoreSystemService(),
		NewCoreProcessService(),
		NewCorePlayersService(),
	)
}
