package steamcm

type ProtocolLayer interface {
	Send([]byte) error
	Handle([]byte) error
	SetOutgoingHandler(func([]byte) error)
	SetIncomingHandler(func([]byte) error)
}
