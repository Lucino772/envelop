package minecraft

import "github.com/Lucino772/envelop/internal/wrapper"

type minecraftModule struct{}

func NewMinecraftModule() *minecraftModule {
	return &minecraftModule{}
}

func (mod *minecraftModule) Register(wrapper wrapper.WrapperRegistrar) {
	wrapper.AddTask(NewCheckMinecraftStatusTask().Run)
	wrapper.AddTask(NewFetchMinecraftPlayersTask().Run)
}
