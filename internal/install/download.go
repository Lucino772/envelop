package install

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Lucino772/envelop/pkg/steam/steamdl"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
)

type DownloaderOptFunc (func(*Downloader))
type Downloader struct {
	httpClient  *http.Client
	steamClient *steamdl.SteamDownloadClient
	config      DownloadConfig
}

func NewDownloader(opts ...DownloaderOptFunc) *Downloader {
	dl := &Downloader{}
	for _, opt := range opts {
		opt(dl)
	}
	return dl
}

func (d *Downloader) Initialize() error {
	if d.steamClient != nil {
		if err := d.steamClient.Connect(); err != nil {
			return err
		}
		if err := d.steamClient.WaitReady(time.Second * 5); err != nil {
			return err
		}
		logInResult, err := d.steamClient.LogInAnonymously()
		if err != nil {
			return err
		}
		log.Println("LogIn", logInResult)
		if logInResult.GetEresult() != int32(steamlang.EResult_OK) {
			return errors.New("failed to intialize downloader")
		}
		return nil
	}
	return nil
}

func (d *Downloader) GetHttpClient() *http.Client {
	return d.httpClient
}

func (d *Downloader) GetSteamClient() *steamdl.SteamDownloadClient {
	return d.steamClient
}

func (d *Downloader) GetConfig() DownloadConfig {
	return d.config
}

func WithHttpClient() DownloaderOptFunc {
	return func(d *Downloader) {
		if d.httpClient == nil {
			d.httpClient = &http.Client{}
		}
	}
}

func WithSteamClient() DownloaderOptFunc {
	return func(d *Downloader) {
		if d.steamClient == nil {
			d.steamClient = steamdl.NewSteamDownloadClient()
		}
	}
}

func WithDownloadConfig(config DownloadConfig) DownloaderOptFunc {
	return func(d *Downloader) {
		d.config = config
	}
}
