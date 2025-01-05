package steamcm

import (
	"errors"
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"google.golang.org/protobuf/proto"
)

type SteamUnifiedMessageHandler struct{}

type steamUnifiedMessageResponse struct {
	Result     steamlang.EResult
	MethodName string
	Packet     *Packet
}

func NewSteamUnifiedMessageHandler() *SteamUnifiedMessageHandler {
	return &SteamUnifiedMessageHandler{}
}

func (handler *SteamUnifiedMessageHandler) Register(handlers map[steamlang.EMsg]func(*Packet) ([]Event, error)) {
	handlers[steamlang.EMsg_ServiceMethodResponse] = handler.handleServiceMethodResponse
	handlers[steamlang.EMsg_ServiceMethod] = handler.handleServiceMethod
}

func (handler *SteamUnifiedMessageHandler) SendMessage(conn Connection, name string, body any) (*steamUnifiedMessageResponse, error) {
	// TODO : Check that user is logged-in
	jobId := conn.GetNextJobId()

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ServiceMethodCallFromClient)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.JobidSource = proto.Uint64(uint64(jobId))
	header.Proto.TargetJobName = proto.String(name)
	encoder.Body = body

	packet, err := encoder.Encode()
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}
	return waitForJob[*steamUnifiedMessageResponse](conn, jobId, time.Second*30)
}

func (handler *SteamUnifiedMessageHandler) SendNotification(conn Connection, name string, body any) error {
	// TODO : Check that user is logged-in

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ServiceMethodCallFromClient)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.TargetJobName = proto.String(name)
	encoder.Body = body

	packet, err := encoder.Encode()
	if err != nil {
		return err
	}
	return conn.SendPacket(packet)
}

func (handler *SteamUnifiedMessageHandler) handleServiceMethodResponse(packet *Packet) ([]Event, error) {
	if !packet.IsProto() {
		return nil, errors.New("non-protobuf packet")
	}
	protoHeader := packet.Header().(*ProtoHeader).Proto
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId: steam.JobId(packet.header.GetTargetJobId()),
			Payload: &steamUnifiedMessageResponse{
				Result:     steamlang.EResult(protoHeader.GetEresult()),
				MethodName: protoHeader.GetTargetJobName(),
				Packet:     packet,
			},
		}),
	}, nil
}

func (handler *SteamUnifiedMessageHandler) handleServiceMethod(packet *Packet) ([]Event, error) {
	if !packet.IsProto() {
		return nil, errors.New("non-protobuf packet")
	}
	// TODO : Implement this
	// protoHeader := packet.Header().(*ProtoHeader).Proto
	return nil, nil
}
