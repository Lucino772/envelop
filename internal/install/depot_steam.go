package install

import (
	"bytes"
	"context"
	"crypto/sha1"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Lucino772/envelop/pkg/steam/steamcdn"
	"github.com/Lucino772/envelop/pkg/steam/steamdl"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steamvdf"
	"github.com/alitto/pond/v2"
)

type SteamDepotManifest struct {
	AppId uint32 `mapstructure:"appid,omitempty"`
}

func (manifest *SteamDepotManifest) GetDownloaderOptions() []DownloaderOptFunc {
	return []DownloaderOptFunc{WithSteamClient()}
}

func (manifest *SteamDepotManifest) GetMetadata(ctx context.Context, dl *Downloader, path string) (Metadata, error) {
	client := dl.GetSteamClient()

	appInfo, err := client.GetApplicationInfo(manifest.AppId)
	if err != nil {
		return nil, err
	}

	kv, err := steamvdf.ReadBytes(appInfo.GetBuffer())
	if err != nil {
		return nil, err
	}

	depots, _ := kv.GetChild("depots")
	depotIds := make([]int, 0)
	for _, depot := range depots.Children {
		depotId, err := strconv.Atoi(depot.Key)
		if err != nil {
			continue
		}
		if len(depot.Children) == 0 {
			continue
		}

		if config, ok := depot.GetChild("config"); ok {
			if oslist, ok := config.ToMapInner()["oslist"]; ok {
				if oslist == dl.GetConfig().TargetOs {
					depotIds = append(depotIds, depotId)
				}
			} else {
				depotIds = append(depotIds, depotId)
			}
		} else {
			depotIds = append(depotIds, depotId)
		}
	}

	depotsMetadatas := make([]steamSourceDepotMetadata, 0)
	for _, depotId := range depotIds {
		depotInfo, err := client.GetDepotInfo(appInfo, uint32(depotId), "public")
		if err != nil {
			continue
		}
		depotManifest, err := client.DownloadDepotManifest(depotInfo)
		if err != nil {
			return nil, err
		}
		depotsMetadatas = append(depotsMetadatas, steamSourceDepotMetadata{
			depotInfo:     depotInfo,
			depotManifest: depotManifest,
		})
	}

	return &SteamSourceMetadata{
		AppId:       manifest.AppId,
		Destination: path,
		Depots:      depotsMetadatas,
	}, nil
}

type SteamSourceMetadata struct {
	AppId       uint32
	Destination string
	Depots      []steamSourceDepotMetadata
}

type steamSourceDepotMetadata struct {
	depotInfo     *steamdl.DepotInfo
	depotManifest *steamcdn.DepotManifest
}

func (metadata *SteamSourceMetadata) Install(ctx context.Context, pool pond.Pool, dl *Downloader) (Waiter, error) {
	rootGroup := pool.NewGroup()
	for _, depot := range metadata.Depots {
		for _, file := range depot.depotManifest.Files {
			// If file is a directory, create and return
			path := filepath.FromSlash(filepath.Join(metadata.Destination, file.Filename))
			if steamlang.EDepotFileFlag_Directory&file.Flags != 0 {
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					return nil, err
				}
			} else {
				// If normal file, create the file and pre-allocate the space
				dir := filepath.Dir(path)
				if err := os.MkdirAll(dir, os.ModePerm); err != nil {
					return nil, err
				}

				fp, err := os.Create(path)
				if err != nil {
					return nil, err
				}
				if err := fp.Truncate(int64(file.TotalSize)); err != nil {
					fp.Close()
					return nil, err
				}
				fp.Close()

				// Submit chunk download
				chunkGroup := pool.NewGroup()
				for _, chunk := range file.Chunks {
					chunkGroup.SubmitErr(func() error {
						// Download chunk
						chunkData, err := dl.GetSteamClient().DownloadDepotChunk(
							depot.depotInfo,
							chunk,
						)
						if err != nil {
							return err
						}

						// Write to file
						_fp, err := os.OpenFile(path, os.O_WRONLY, 0)
						if err != nil {
							return err
						}
						defer _fp.Close()
						if _, err := _fp.Seek(int64(chunk.Offset), 0); err != nil {
							return err
						}
						if _, err := _fp.Write(chunkData.Data); err != nil {
							return err
						}
						return nil
					})
				}

				// Submit file cleanup task
				rootGroup.SubmitErr(func() error {
					groupErr := chunkGroup.Wait()
					if groupErr != nil {
						return groupErr
					}

					// Check file hash
					fpr, err := os.Open(path)
					if err != nil {
						return err
					}
					hasher := sha1.New()
					if _, err := io.Copy(hasher, fpr); err != nil {
						return err
					}
					if value := hasher.Sum(nil); !bytes.Equal(value, file.FileHash) {
						return errors.New("hash mismatch")
					}
					return nil
				})
			}
		}
	}
	return rootGroup, nil
}
