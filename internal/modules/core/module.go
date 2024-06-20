package core

import "github.com/Lucino772/envelop/internal/wrapper"

type coreModule struct{}

func NewCoreModule() *coreModule {
	return &coreModule{}
}

func (mod *coreModule) Register(wrapper wrapper.WrapperRegistrar) {
	wrapper.AddService(NewCoreSystemService())
	wrapper.AddService(NewCoreProcessService())
	wrapper.AddService(NewCorePlayersService())
}
