package install

import "net/http"

type DownloaderOptFunc (func(*Downloader))
type Downloader struct {
	httpClient *http.Client
	config     DownloadConfig
}

func NewDownloader(opts ...DownloaderOptFunc) *Downloader {
	dl := &Downloader{}
	for _, opt := range opts {
		opt(dl)
	}
	return dl
}

func (d *Downloader) GetHttpClient() *http.Client {
	return d.httpClient
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

func WithDownloadConfig(config DownloadConfig) DownloaderOptFunc {
	return func(d *Downloader) {
		d.config = config
	}
}
