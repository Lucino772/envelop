package steamcm

import (
	"fmt"

	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steammsg"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"google.golang.org/protobuf/proto"
)

type SteamContentHandler struct {
	unifiedMessage *SteamUnifiedMessageHandler
}

func NewSteamContentHandler(unifiedMessage *SteamUnifiedMessageHandler) *SteamContentHandler {
	return &SteamContentHandler{unifiedMessage: unifiedMessage}
}

func (handler *SteamContentHandler) Register(_ map[steamlang.EMsg]func(*steammsg.Packet) ([]Event, error)) {
}

func (handler *SteamContentHandler) GetManifestRequestCode(conn Connection, depotId uint32, appId uint32, manifestId uint64, branch string) (uint64, error) {
	request := &steampb.CContentServerDirectory_GetManifestRequestCode_Request{
		AppId:              proto.Uint32(appId),
		DepotId:            proto.Uint32(depotId),
		ManifestId:         proto.Uint64(manifestId),
		AppBranch:          proto.String(branch),
		BranchPasswordHash: nil,
	}
	response, err := handler.unifiedMessage.SendMessage(
		conn,
		fmt.Sprintf("%s.%s#%s", "ContentServerDirectory", "GetManifestRequestCode", "1"),
		request,
	)
	if err != nil {
		return 0, err
	}

	body := new(steampb.CContentServerDirectory_GetManifestRequestCode_Response)
	if _, err := steammsg.DecodePacket(response.Packet, body); err != nil {
		return 0, err
	}
	return body.GetManifestRequestCode(), nil
}

func (handler *SteamContentHandler) GetServersForSteamPipe(conn Connection, cellId uint32) ([]*steampb.CContentServerDirectory_ServerInfo, error) {
	request := &steampb.CContentServerDirectory_GetServersForSteamPipe_Request{CellId: &cellId}
	response, err := handler.unifiedMessage.SendMessage(
		conn,
		fmt.Sprintf("%s.%s#%s", "ContentServerDirectory", "GetServersForSteamPipe", "1"),
		request,
	)
	if err != nil {
		return nil, err
	}

	body := new(steampb.CContentServerDirectory_GetServersForSteamPipe_Response)
	if _, err := steammsg.DecodePacket(response.Packet, body); err != nil {
		return nil, err
	}
	return body.GetServers(), nil
}

func (handler *SteamContentHandler) GetCDNAuthToken(conn Connection, appId uint32, depotId uint32, cdnHostname string) (string, error) {
	request := &steampb.CContentServerDirectory_GetCDNAuthToken_Request{
		AppId:    proto.Uint32(appId),
		DepotId:  proto.Uint32(depotId),
		HostName: proto.String(cdnHostname),
	}
	response, err := handler.unifiedMessage.SendMessage(
		conn,
		fmt.Sprintf("%s.%s#%s", "ContentServerDirectory", "GetCDNAuthToken", "1"),
		request,
	)
	if err != nil {
		return "", err
	}

	body := new(steampb.CContentServerDirectory_GetCDNAuthToken_Response)
	if _, err := steammsg.DecodePacket(response.Packet, body); err != nil {
		return "", err
	}
	return body.GetToken(), nil
}
