package install

import (
	"context"

	"github.com/alitto/pond/v2"
)

type DownloadContext struct {
	OsName      string
	OsArch      string
	OsLang      string
	InstallDir  string
	LowViolence bool
}

type Manifest struct {
	Sources []Source `mapstructure:"sources,omitempty"`
	Config  string   `mapstructure:"config,omitempty"`
}

type Source interface {
	GetDownloaderOptions() []DownloaderOptFunc
	GetMetadata(context.Context, DownloadContext, *Downloader) (Metadata, error)
}

type Metadata interface {
	GetExports() map[string]any
	Install(context.Context, pond.Pool, *Downloader) (Waiter, error)
}

type Waiter interface {
	Done() <-chan struct{}
	Wait() error
}
