package steamdl

import (
	"errors"
	"slices"
	"strconv"
	"time"

	"github.com/Lucino772/envelop/pkg/steam/steamcm"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"github.com/Lucino772/envelop/pkg/steam/steamvdf"
)

type DepotInfo struct {
	AppId         uint32
	DepotId       uint32
	ManifestId    uint64
	Branch        string
	DecryptionKey []byte
}

type SteamDownloadClient struct {
	conn *steamcm.SteamConnection

	// Handlers
	user    *steamcm.SteamUserHandler
	apps    *steamcm.SteamAppsHandler
	content *steamcm.SteamContentHandler
}

func NewSteamDownloadClient() *SteamDownloadClient {
	user := steamcm.NewUserHandler()
	apps := steamcm.NewAppsHandler()
	unified := steamcm.NewSteamUnifiedMessageHandler()
	content := steamcm.NewSteamContentHandler(unified)

	conn := steamcm.NewSteamConnection(
		steamcm.NewSteamBaseHandler(),
		user,
		apps,
		unified,
		content,
	)
	return &SteamDownloadClient{
		conn:    conn,
		user:    user,
		apps:    apps,
		content: content,
	}
}

func (client *SteamDownloadClient) WaitReady(timeout time.Duration) error {
	return client.conn.WaitReady(timeout)
}

func (client *SteamDownloadClient) Connect() error {
	return client.conn.Connect()
}

func (client *SteamDownloadClient) LogInAnonymously() (*steampb.CMsgClientLogonResponse, error) {
	return client.user.LogInAnonymously(client.conn)
}

func (client *SteamDownloadClient) GetApplicationInfo(appId uint32) (*steampb.CMsgClientPICSProductInfoResponse_AppInfo, error) {
	accessTokens, err := client.apps.PICSGetAccessTokens(
		client.conn,
		[]steamcm.PICSRequest{{ID: appId, AccessToken: 0}},
		[]steamcm.PICSRequest{},
	)
	if err != nil {
		return nil, err
	}
	if slices.Contains(accessTokens.AppDeniedTokens, appId) {
		return nil, errors.New("access denied")
	}

	requests := make([]steamcm.PICSRequest, 0)
	for _, app := range accessTokens.GetAppAccessTokens() {
		requests = append(requests, steamcm.PICSRequest{
			ID:          app.GetAppid(),
			AccessToken: app.GetAccessToken(),
		})
	}

	appInfos, err := client.apps.PICSGetProductInfo(
		client.conn,
		requests,
		[]steamcm.PICSRequest{},
		false,
	)
	if err != nil {
		return nil, err
	}

	appInfo := appInfos.Apps[0]
	kv, err := steamvdf.ReadBytes(appInfo.GetBuffer())
	if err != nil {
		return nil, err
	}

	hasAccess, err := client.accountHasAccess(appId, appId)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		if common, ok := kv.GetChild("common"); ok {
			if freeToDownload, ok := common.GetChild("freetodownload"); ok {
				hasAccess = freeToDownload.Value == "1"
			}
		}
	}
	if !hasAccess {
		return nil, errors.New("no access to app")
	}

	return appInfo, nil
}

func (client *SteamDownloadClient) GetDepotInfo(appInfo *steampb.CMsgClientPICSProductInfoResponse_AppInfo, depotId uint32, branch string) (*DepotInfo, error) {
	manifestId, appId, err := client.getDepotManifestId(appInfo, depotId, branch)
	if err != nil {
		return nil, err
	}
	if manifestId == 0 {
		return nil, errors.New("invalid manifest id")
	}

	depotKeyResponse, err := client.apps.GetDepotDecryptionKey(client.conn, depotId, appId)
	if err != nil {
		return nil, err
	}

	return &DepotInfo{
		AppId:         appId,
		DepotId:       depotId,
		ManifestId:    manifestId,
		Branch:        branch,
		DecryptionKey: depotKeyResponse.GetDepotEncryptionKey(),
	}, nil
}

func (client *SteamDownloadClient) DownloadDepotManifest(depotInfo *DepotInfo) (*steamcm.DepotManifest, error) {
	manifestRequestCode, err := client.content.GetManifestRequestCode(
		client.conn,
		depotInfo.DepotId,
		depotInfo.AppId,
		depotInfo.ManifestId,
		depotInfo.Branch,
	)
	if err != nil {
		return nil, err
	}
	return steamcm.NewCDNClient().DownloadManifest(
		"fastly.cdn.steampipe.steamcontent.com",
		depotInfo.DepotId,
		depotInfo.ManifestId,
		manifestRequestCode,
		depotInfo.DecryptionKey,
	)
}

func (client *SteamDownloadClient) GetCDNAuthToken(appId uint32, depotId uint32) (string, error) {
	return client.content.GetCDNAuthToken(
		client.conn,
		appId,
		depotId,
		"fastly.cdn.steampipe.steamcontent.com",
	)
}

func (client *SteamDownloadClient) DownloadDepotChunk(depotInfo *DepotInfo, chunk steamcm.ChunkData, cdnAuthToken string) (*steamcm.DepotChunk, error) {
	return steamcm.NewCDNClient().DownloadDepotChunk(
		"fastly.cdn.steampipe.steamcontent.com",
		depotInfo.DepotId,
		chunk,
		depotInfo.DecryptionKey,
		cdnAuthToken,
	)
}

func (client *SteamDownloadClient) getDepotManifestId(appInfo *steampb.CMsgClientPICSProductInfoResponse_AppInfo, depotId uint32, branch string) (uint64, uint32, error) {
	kv, err := steamvdf.ReadBytes(appInfo.GetBuffer())
	if err != nil {
		return 0, 0, err
	}

	depots, ok := kv.GetChild("depots")
	if !ok {
		return 0, 0, nil
	}
	depot, ok := depots.GetChild(strconv.FormatUint(uint64(depotId), 10))
	if !ok {
		return 0, 0, nil
	}

	manifests, hasManifests := depot.GetChild("manifests")
	depotFromApp, hasDepotFromApp := depot.GetChild("depotfromapp")

	if !hasManifests && hasDepotFromApp {
		otherAppId, err := strconv.ParseUint(depotFromApp.Value, 10, 0)
		if err != nil {
			return 0, 0, err
		}
		if otherAppId == uint64(appInfo.GetAppid()) {
			return 0, 0, nil
		}
		otherAppInfo, err := client.GetApplicationInfo(uint32(otherAppId))
		if err != nil {
			return 0, 0, err
		}
		return client.getDepotManifestId(otherAppInfo, depotId, branch)
	}

	// TODO: Add support for encrypted manifests
	if manifest, ok := manifests.GetChild(branch); ok {
		if gid, ok := manifest.GetChild("gid"); ok {
			manifestId, err := strconv.ParseUint(gid.Value, 10, 0)
			if err != nil {
				return 0, 0, err
			}
			return manifestId, appInfo.GetAppid(), nil
		}
	}

	return 0, 0, nil
}

func (client *SteamDownloadClient) accountHasAccess(appId uint32, depotId uint32) (bool, error) {
	licenseInfo, err := client.apps.PICSGetProductInfo(
		client.conn,
		[]steamcm.PICSRequest{},
		[]steamcm.PICSRequest{{ID: 17906}},
		false,
	)
	if err != nil {
		return false, err
	}

	for _, license := range licenseInfo.Packages {
		kv, err := steamvdf.ReadBytes(license.GetBuffer())
		if err != nil {
			return false, err
		}
		if appIds, ok := kv.GetChild("appids"); ok {
			if slices.Contains(appIds.GetChildrenAsSlice(), strconv.FormatUint(uint64(appId), 10)) {
				return true, nil
			}
		}
		if depotIds, ok := kv.GetChild("depotids"); ok {
			if slices.Contains(depotIds.GetChildrenAsSlice(), strconv.FormatUint(uint64(depotId), 10)) {
				return true, nil
			}
		}
	}

	return false, nil
}
