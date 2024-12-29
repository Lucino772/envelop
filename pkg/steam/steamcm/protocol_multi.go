package steamcm

type multiLayer struct {
	layers []ProtocolLayer
}

func NewMultiLayer(layers ...ProtocolLayer) *multiLayer {
	layer := &multiLayer{layers: layers}
	for i := 0; i < len(layer.layers); i++ {
		if i < len(layer.layers)-1 {
			layer.layers[i].SetIncomingHandler(layer.layers[i+1].Handle)
		}
		if i > 0 {
			layer.layers[i].SetOutgoingHandler(layer.layers[i-1].Send)
		}
	}
	return layer
}

func (layer *multiLayer) Send(data []byte) error {
	return layer.layers[len(layer.layers)-1].Send(data)
}

func (layer *multiLayer) Handle(data []byte) error {
	return layer.layers[0].Handle(data)
}

func (layer *multiLayer) SetOutgoingHandler(handler func([]byte) error) {
	layer.layers[0].SetOutgoingHandler(handler)
}

func (layer *multiLayer) SetIncomingHandler(handler func([]byte) error) {
	layer.layers[len(layer.layers)-1].SetIncomingHandler(handler)
}
