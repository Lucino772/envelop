package modules

import (
	"github.com/Lucino772/envelop/internal/modules/configs"
	"github.com/Lucino772/envelop/internal/modules/core"
	"github.com/Lucino772/envelop/internal/modules/minecraft"
	"github.com/Lucino772/envelop/internal/modules/rcon"
	"github.com/Lucino772/envelop/internal/wrapper"
)

func InitializeModule(name string, opts map[string]any, registry *wrapper.Registry) {
	switch name {
	case "envelop.core":
		core.Initialize(opts, registry)
	case "envelop.configs":
		configs.Initialize(opts, registry)
	case "envelop.rcon":
		rcon.Initialize(opts, registry)
	case "envelop.minecraft":
		minecraft.Initialize(opts, registry)
	}
}
