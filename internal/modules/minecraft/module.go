package minecraft

import "github.com/Lucino772/envelop/internal/wrapper"

type minecraftModule struct{}

func NewMinecraftModule() *minecraftModule {
	return &minecraftModule{}
}

func (mod *minecraftModule) Register(w wrapper.Wrapper) []wrapper.OptFunc {
	return []wrapper.OptFunc{
		wrapper.WithService(NewMinecraftRconService(w)),
		wrapper.WithTask(NewCheckMinecraftStatusTask().Run),
		wrapper.WithTask(NewFetchMinecraftPlayersTask().Run),
	}
}
