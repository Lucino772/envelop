package steamcm

import (
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steammsg"
)

type Handler interface {
	Register(map[steamlang.EMsg]func(*steammsg.Packet) ([]Event, error))
}

func MakeDispatchLayer(handlers ...Handler) *dispatchLayer {
	dispatchMap := make(map[steamlang.EMsg]func(*steammsg.Packet) ([]Event, error), 0)
	for _, handler := range handlers {
		handler.Register(dispatchMap)
	}
	return NewDispatchLayer(dispatchMap)
}
