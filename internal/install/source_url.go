package install

import (
	"path/filepath"

	"github.com/Lucino772/envelop/pkg/download"
)

type UrlSource struct {
	Type        string         `mapstructure:"type,omitempty"`
	Destination string         `mapstructure:"destination,omitempty"`
	Exports     map[string]any `mapstructure:"exports,omitempty"`
	Url         string         `mapstructure:"url,omitempty"`
}

func (s *UrlSource) WithInstallDir(dir string) Source {
	return &UrlSource{
		Type:        s.Type,
		Destination: filepath.Join(dir, s.Destination),
		Exports:     s.Exports,
		Url:         s.Url,
	}
}

func (s *UrlSource) GetExports() map[string]any {
	data := struct{ Destination string }{
		Destination: s.Destination,
	}
	return parseExports(s.Exports, data)
}

func (s *UrlSource) IterDownloaders(yield func(Downloader) bool) {
	downloader := download.NewDownloader(s.Url, s.Destination)
	if !yield(downloader) {
		return
	}
}
