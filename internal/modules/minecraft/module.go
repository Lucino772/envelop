package minecraft

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any) wrapper.Module {
	return func(w wrapper.Wrapper) []wrapper.OptFunc {
		return []wrapper.OptFunc{
			wrapper.WithService(NewMinecraftRconService(w)),
			wrapper.WithTask(NewCheckMinecraftStatusTask()),
			wrapper.WithTask(NewFetchMinecraftPlayersTask()),
		}
	}
}
