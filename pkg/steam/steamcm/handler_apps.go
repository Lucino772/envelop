package steamcm

import (
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type SteamAppsHandler struct{}

type PICSRequest struct {
	ID          uint32
	AccessToken uint64
}

func NewAppsHandler() *SteamAppsHandler {
	return &SteamAppsHandler{}
}

func (handler *SteamAppsHandler) Register(handlers map[steamlang.EMsg]func(*Packet) ([]Event, error)) {
	handlers[steamlang.EMsg_ClientPICSProductInfoResponse] = handler.handlePICSProductInfoResponse
	handlers[steamlang.EMsg_ClientPICSAccessTokenResponse] = handler.handlePICSGetAccessTokensResponse
	handlers[steamlang.EMsg_ClientRequestFreeLicenseResponse] = handler.handleFreeLicenseResponse
	handlers[steamlang.EMsg_ClientGetDepotDecryptionKeyResponse] = handler.handleGetDepotDecryptionKeyResponse
	handlers[steamlang.EMsg_ClientGetCDNAuthTokenResponse] = handler.handleGetCDNAuthTokenResponse
}

func (handler *SteamAppsHandler) PICSGetAccessTokens(conn Connection, apps []PICSRequest, packages []PICSRequest) (*steampb.CMsgClientPICSAccessTokenResponse, error) {
	jobId := conn.GetNextJobId()

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ClientPICSAccessTokenRequest)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.JobidSource = proto.Uint64(uint64(jobId))
	var body = &steampb.CMsgClientPICSAccessTokenRequest{}

	for _, app := range apps {
		body.Appids = append(body.Appids, app.ID)
	}
	for _, pkg := range packages {
		body.Packageids = append(body.Packageids, pkg.ID)
	}
	encoder.Body = body

	packet, err := encoder.Encode()
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}
	return waitForJob[*steampb.CMsgClientPICSAccessTokenResponse](conn, jobId, time.Second*30)
}

func (handler *SteamAppsHandler) PICSGetProductInfo(conn Connection, apps []PICSRequest, packages []PICSRequest, onlyMetaData bool) (*steampb.CMsgClientPICSProductInfoResponse, error) {
	jobId := conn.GetNextJobId()

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ClientPICSProductInfoRequest)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.JobidSource = proto.Uint64(uint64(jobId))

	var body = &steampb.CMsgClientPICSProductInfoRequest{}
	for _, app := range apps {
		body.Apps = append(body.Apps, &steampb.CMsgClientPICSProductInfoRequest_AppInfo{
			AccessToken:        proto.Uint64(app.AccessToken),
			Appid:              proto.Uint32(app.ID),
			OnlyPublicObsolete: proto.Bool(false),
		})
	}

	for _, pkg := range packages {
		body.Packages = append(body.Packages, &steampb.CMsgClientPICSProductInfoRequest_PackageInfo{
			AccessToken: proto.Uint64(pkg.AccessToken),
			Packageid:   proto.Uint32(pkg.ID),
		})
	}

	body.MetaDataOnly = proto.Bool(onlyMetaData)
	encoder.Body = body

	packet, err := encoder.Encode()
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}
	return waitForJob[*steampb.CMsgClientPICSProductInfoResponse](conn, jobId, time.Second*30)
}

func (handler *SteamAppsHandler) RequestFreeLicense(conn Connection, appIds []uint32) (*steampb.CMsgClientRequestFreeLicenseResponse, error) {
	jobId := conn.GetNextJobId()

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ClientRequestFreeLicense)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.JobidSource = proto.Uint64(uint64(jobId))
	encoder.Body = &steampb.CMsgClientRequestFreeLicense{
		Appids: appIds,
	}
	packet, err := encoder.Encode()
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}
	return waitForJob[*steampb.CMsgClientRequestFreeLicenseResponse](conn, jobId, time.Second*30)
}

func (handler *SteamAppsHandler) GetDepotDecryptionKey(conn Connection, depotId uint32, appId uint32) (*steampb.CMsgClientGetDepotDecryptionKeyResponse, error) {
	jobId := conn.GetNextJobId()

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ClientGetDepotDecryptionKey)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.JobidSource = proto.Uint64(uint64(jobId))
	encoder.Body = &steampb.CMsgClientGetDepotDecryptionKey{
		DepotId: proto.Uint32(depotId),
		AppId:   proto.Uint32(appId),
	}
	packet, err := encoder.Encode()
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}
	return waitForJob[*steampb.CMsgClientGetDepotDecryptionKeyResponse](conn, jobId, time.Second*30)
}

func (handler *SteamAppsHandler) GetCDNAuthToken(conn Connection, appId uint32, depotId uint32, serverHostname string) (*steampb.CMsgClientGetCDNAuthTokenResponse, error) {
	jobId := conn.GetNextJobId()

	var encoder = NewProtoPacketEncoder(steamlang.EMsg_ClientGetCDNAuthToken)
	header := encoder.Header.(*ProtoHeader)
	header.Proto.JobidSource = proto.Uint64(uint64(jobId))
	encoder.Body = &steampb.CMsgClientGetCDNAuthToken{
		AppId:    proto.Uint32(appId),
		DepotId:  proto.Uint32(depotId),
		HostName: proto.String(serverHostname),
	}
	packet, err := encoder.Encode()
	if err != nil {
		return nil, err
	}
	if err := conn.SendPacket(packet); err != nil {
		return nil, err
	}
	return waitForJob[*steampb.CMsgClientGetCDNAuthTokenResponse](conn, jobId, time.Second*30)
}

func (handler *SteamAppsHandler) handlePICSGetAccessTokensResponse(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientPICSAccessTokenResponse]{
		Body: new(steampb.CMsgClientPICSAccessTokenResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.header.GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}

func (handler *SteamAppsHandler) handlePICSProductInfoResponse(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientPICSProductInfoResponse]{
		Body: new(steampb.CMsgClientPICSProductInfoResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.header.GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}

func (handler *SteamAppsHandler) handleFreeLicenseResponse(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientRequestFreeLicenseResponse]{
		Body: new(steampb.CMsgClientRequestFreeLicenseResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.header.GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}

func (handler *SteamAppsHandler) handleGetDepotDecryptionKeyResponse(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientGetDepotDecryptionKeyResponse]{
		Body: new(steampb.CMsgClientGetDepotDecryptionKeyResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.header.GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}

func (handler *SteamAppsHandler) handleGetCDNAuthTokenResponse(packet *Packet) ([]Event, error) {
	var decoder = &ProtoPacketDecoder[*steampb.CMsgClientGetCDNAuthTokenResponse]{
		Body: new(steampb.CMsgClientGetCDNAuthTokenResponse),
	}
	if err := decoder.Decode(packet); err != nil {
		return nil, err
	}
	return []Event{
		MakeEvent(EventType_State, EventCallback{
			JobId:   steam.JobId(packet.header.GetTargetJobId()),
			Payload: decoder.Body,
		}),
	}, nil
}
