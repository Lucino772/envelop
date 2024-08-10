package core

import "github.com/Lucino772/envelop/internal/wrapper"

type coreModule struct{}

func NewCoreModule() *coreModule {
	return &coreModule{}
}

func (mod *coreModule) Register(w wrapper.Wrapper) []wrapper.OptFunc {
	return []wrapper.OptFunc{
		wrapper.WithService(NewCoreSystemService(w)),
		wrapper.WithService(NewCoreProcessService(w)),
		wrapper.WithService(NewCorePlayersService(w)),
	}
}
