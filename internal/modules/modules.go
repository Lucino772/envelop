package modules

import (
	"github.com/Lucino772/envelop/internal/modules/core"
	"github.com/Lucino772/envelop/internal/modules/minecraft"
	"github.com/Lucino772/envelop/internal/wrapper"
)

func InitializeModule(name string, opts map[string]any, wrapperOpts *wrapper.Options) {
	switch name {
	case "envelop.core":
		core.Initialize(opts, wrapperOpts)
	case "envelop.minecraft":
		minecraft.Initialize(opts, wrapperOpts)
	}
}
