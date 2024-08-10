package modules

import (
	"github.com/Lucino772/envelop/internal/modules/core"
	"github.com/Lucino772/envelop/internal/modules/minecraft"
	"github.com/Lucino772/envelop/internal/wrapper"
)

func NewModule(name string, opts map[string]any) wrapper.Module {
	switch name {
	case "envelop.core":
		return core.Initialize(opts)
	case "envelop.minecraft":
		return minecraft.Initialize(opts)
	default:
		return nil
	}
}
