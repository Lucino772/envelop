package install

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Lucino772/envelop/pkg/download"
	"github.com/alitto/pond/v2"
)

type HttpSource struct {
	Type        string         `mapstructure:"type,omitempty"`
	Destination string         `mapstructure:"destination,omitempty"`
	Exports     map[string]any `mapstructure:"exports,omitempty"`
	Url         string         `mapstructure:"url,omitempty"`
	Hash        string         `mapstructure:"hash,omitempty"`
	Archive     string         `mapstructure:"archive,omitempty"`
}

func (s *HttpSource) GetDownloaderOptions() []DownloaderOptFunc {
	return []DownloaderOptFunc{WithHttpClient()}
}

func (s *HttpSource) GetMetadata(ctx context.Context, dl *Downloader) (Metadata, error) {
	var (
		decompressor download.Decompressor
		checksum     *download.Checksum
	)

	contentLen, acceptRanges, err := s.fetchDownloadInfo(ctx, dl.GetHttpClient())
	if err != nil {
		return nil, err
	}

	if s.Archive != "" {
		switch s.Archive {
		case "zip":
			decompressor = &download.ZipDecompressor{}
		case "tar":
			decompressor = &download.TarDecompressor{}
		case "tar:gz":
			decompressor = &download.TarGzipDecompressor{}
		default:
			return nil, download.ErrArchiveFormatNotSupported
		}
	}

	if s.Hash != "" {
		values := strings.SplitN(s.Hash, ":", 2)
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

	return &HttpSourceMetadata{
		Url:          s.Url,
		Destination:  filepath.Join(dl.GetConfig().InstallDir, s.Destination),
		Exports:      s.Exports,
		ContentLen:   contentLen,
		AcceptRanges: acceptRanges,
		Decompressor: decompressor,
		Checksum:     checksum,
	}, nil
}

func (s *HttpSource) fetchDownloadInfo(ctx context.Context, client *http.Client) (int64, bool, error) {
	request, err := http.NewRequestWithContext(ctx, "HEAD", s.Url, nil)
	if err != nil {
		return -1, false, err
	}
	response, err := client.Do(request)
	if err != nil {
		return -1, false, err
	}
	if err := response.Body.Close(); err != nil {
		return -1, false, err
	}
	if response.StatusCode != 200 {
		return -1, false, err
	}
	return response.ContentLength, response.Header.Get("Accept-Ranges") == "bytes", nil
}

type HttpSourceMetadata struct {
	Url          string
	Destination  string
	Exports      map[string]any
	ContentLen   int64
	AcceptRanges bool
	Decompressor download.Decompressor
	Checksum     *download.Checksum
}

func (metadata *HttpSourceMetadata) GetExports() map[string]any {
	data := struct{ Destination string }{
		Destination: metadata.Destination,
	}
	return parseExports(metadata.Exports, data)
}

func (metadata *HttpSourceMetadata) Install(ctx context.Context, pool pond.Pool, dl *Downloader) (Waiter, error) {
	group := pool.NewGroupContext(ctx)
	group.SubmitErr(func() error {
		var dst string = metadata.Destination
		var decompressorDst string

		if metadata.Decompressor != nil {
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

		parsedurl, err := url.Parse(metadata.Url)
		if err != nil {
			return err
		}

		getter := &download.HttpGetter{Client: dl.GetHttpClient()}
		if err := getter.Get(ctx, parsedurl, dst); err != nil {
			return err
		}
		if metadata.Checksum != nil {
			if err := metadata.Checksum.Checksum(dst); err != nil {
				return err
			}
		}
		if metadata.Decompressor != nil {
			if err := metadata.Decompressor.Decompress(ctx, dst, decompressorDst); err != nil {
				return err
			}
		}
		return nil
	})
	return group, nil
}
