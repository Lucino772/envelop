package install

import (
	"context"
	"net/url"
	"strings"

	"github.com/Lucino772/envelop/pkg/download"
)

type UrlSource struct {
	Type        string         `mapstructure:"type,omitempty"`
	Destination string         `mapstructure:"destination,omitempty"`
	Exports     map[string]any `mapstructure:"exports,omitempty"`
	Url         string         `mapstructure:"url,omitempty"`
}

func (s *UrlSource) GetMetadata(ctx context.Context, dlCtx DownloadContext) (Metadata, error) {
	var (
		archive  string
		checksum string
	)

	parsedUrl, err := url.Parse(s.Url)
	if err != nil {
		return nil, err
	}

	query := parsedUrl.Query()
	if query.Has("archive") {
		archive = query.Get("archive")
		query.Del("archive")
		parsedUrl.RawQuery = query.Encode()
	}

	if query.Has("checksum") {
		checksum = query.Get("checksum")
		query.Del("checksum")
		parsedUrl.RawQuery = query.Encode()
	}

	switch parsedUrl.Scheme {
	case "http", "https":
		source := &HttpSource{
			Type:        "http",
			Destination: s.Destination,
			Exports:     s.Exports,
			Url:         parsedUrl.String(),
			Hash:        checksum,
			Archive:     archive,
		}
		return source.GetMetadata(ctx, dlCtx)
	case "data":
		url := strings.Split(parsedUrl.RequestURI(), ";")
		source := &Base64Source{
			Type:        "base64",
			Destination: s.Destination,
			Exports:     s.Exports,
			Content:     strings.Split(url[len(url)-1], ",")[1],
		}
		return source.GetMetadata(ctx, dlCtx)
	default:
		return nil, download.ErrSchemeNotSupported
	}
}
