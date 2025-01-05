package install

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Lucino772/envelop/pkg/download"
	"github.com/alitto/pond/v2"
)

var getters = map[string]download.Getter{
	"http":  &download.HttpGetter{},
	"https": &download.HttpGetter{},
	"data":  &download.DataGetter{},
}
var decompressors = map[string]download.Decompressor{
	"zip":    &download.ZipDecompressor{},
	"tar":    &download.TarDecompressor{},
	"tar:gz": &download.TarGzipDecompressor{},
}

type UrlSource struct {
	Type        string         `mapstructure:"type,omitempty"`
	Destination string         `mapstructure:"destination,omitempty"`
	Exports     map[string]any `mapstructure:"exports,omitempty"`
	Url         string         `mapstructure:"url,omitempty"`
}

func (s *UrlSource) GetMetadata(ctx context.Context, dlCtx DownloadContext) (Metadata, error) {
	// TODO: Filter based on other parameters in download context
	// TODO: Make a head request to get information about the file (hash, content size, etc.)
	var (
		getter       download.Getter
		decompressor download.Decompressor
		checksum     *download.Checksum
	)

	parsedUrl, err := url.Parse(s.Url)
	if err != nil {
		return nil, err
	}
	getter, ok := getters[parsedUrl.Scheme]
	if !ok {
		return nil, download.ErrSchemeNotSupported
	}

	query := parsedUrl.Query()
	if query.Has("archive") {
		archiveFormat := query.Get("archive")
		query.Del("archive")
		parsedUrl.RawQuery = query.Encode()

		decompressor, ok = decompressors[archiveFormat]
		if !ok {
			return nil, download.ErrArchiveFormatNotSupported
		}
	}

	if query.Has("checksum") {
		values := strings.SplitN(query.Get("checksum"), ":", 2)
		query.Del("checksum")
		parsedUrl.RawQuery = query.Encode()

		if len(values) != 2 {
			return nil, download.ErrInvalidChecksumFormat
		}
		checksumAlgo, checksumValue := values[0], values[1]

		checksum = new(download.Checksum)
		checksum.Value, err = hex.DecodeString(checksumValue)
		if err != nil {
			return nil, err
		}

		switch checksumAlgo {
		case "md5":
			checksum.Hash = md5.New()
		case "sha1":
			checksum.Hash = sha1.New()
		case "sha256":
			checksum.Hash = sha256.New()
		case "sha512":
			checksum.Hash = sha512.New()
		default:
			return nil, download.ErrChecksumAlgoNotSupported
		}
	}

	return &urlSourceMetadata{
		Url:          parsedUrl,
		Destination:  filepath.Join(dlCtx.InstallDir, s.Destination),
		Getter:       getter,
		Decompressor: decompressor,
		Checksum:     checksum,
		Exports:      s.Exports,
	}, nil
}

type urlSourceMetadata struct {
	Url          *url.URL
	Destination  string
	Exports      map[string]any
	Getter       download.Getter
	Decompressor download.Decompressor
	Checksum     *download.Checksum
}

func (s *urlSourceMetadata) GetExports() map[string]any {
	data := struct{ Destination string }{
		Destination: s.Destination,
	}
	return parseExports(s.Exports, data)
}

func (s *urlSourceMetadata) Install(ctx context.Context, pool pond.Pool) (Waiter, error) {
	group := pool.NewGroupContext(ctx)
	group.SubmitErr(func() error {
		var dst string = s.Destination
		var decompressorDst string

		if s.Decompressor != nil {
			decompressorDst = dst
			tmpFile, err := os.CreateTemp("", "")
			if err != nil {
				if tmpFile != nil {
					tmpFile.Close()
				}
				return err
			}
			defer os.Remove(tmpFile.Name())
			dst = tmpFile.Name()
			tmpFile.Close()
		}

		if err := s.Getter.Get(ctx, s.Url, dst); err != nil {
			return err
		}
		if s.Checksum != nil {
			if err := s.Checksum.Checksum(dst); err != nil {
				return err
			}
		}
		if s.Decompressor != nil {
			if err := s.Decompressor.Decompress(ctx, dst, decompressorDst); err != nil {
				return err
			}
		}
		return nil
	})
	return group, nil
}
