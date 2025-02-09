package minecraft

import (
	"github.com/Lucino772/envelop/internal/wrapper"
)

func Initialize(_ map[string]any, registry *wrapper.Registry) {
	registry.Tasks = append(
		registry.Tasks,
		NewCheckMinecraftStatusTask(),
		NewFetchMinecraftPlayersTask(),
	)
}
