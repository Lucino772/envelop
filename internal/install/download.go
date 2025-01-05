package install

import "net/http"

type DownloaderOptFunc (func(*Downloader))
type Downloader struct {
	httpClient *http.Client
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

func WithHttpClient() DownloaderOptFunc {
	return func(d *Downloader) {
		if d.httpClient == nil {
			d.httpClient = &http.Client{}
		}
	}
}
