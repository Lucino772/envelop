package minecraft

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any, wrapperOpts *wrapper.Options) {
	wrapperOpts.Services = append(
		wrapperOpts.Services,
		NewMinecraftRconService(),
	)
	wrapperOpts.Tasks = append(
		wrapperOpts.Tasks,
		NewCheckMinecraftStatusTask(),
		NewFetchMinecraftPlayersTask(),
	)
}
