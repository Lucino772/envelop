package core

import "github.com/Lucino772/envelop/internal/wrapper"

func Initialize(_ map[string]any) wrapper.Module {
	return func(w wrapper.Wrapper) []wrapper.OptFunc {
		return []wrapper.OptFunc{
			wrapper.WithService(NewCoreSystemService(w)),
			wrapper.WithService(NewCoreProcessService(w)),
			wrapper.WithService(NewCorePlayersService(w)),
		}
	}
}
