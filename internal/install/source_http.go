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
	"strings"

	"github.com/Lucino772/envelop/pkg/download"
)

type HttpSource struct {
	Url     string `mapstructure:"url,omitempty"`
	Hash    string `mapstructure:"hash,omitempty"`
	Archive string `mapstructure:"archive,omitempty"`
}

func (s *HttpSource) GetDownloaderOptions() []DownloaderOptFunc {
	return []DownloaderOptFunc{WithHttpClient()}
}

func (s *HttpSource) Download(ctx context.Context, dl *Downloader, dst string) error {
	var (
		decompressor download.Decompressor
		checksum     *download.Checksum
	)

	// Prepare download
	_, _, err := s.fetchDownloadInfo(ctx, dl.GetHttpClient())
	if err != nil {
		return err
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
			return download.ErrArchiveFormatNotSupported
		}
	}

	if s.Hash != "" {
		values := strings.SplitN(s.Hash, ":", 2)
		if len(values) != 2 {
			return download.ErrInvalidChecksumFormat
		}
		checksumAlgo, checksumValue := values[0], values[1]

		checksum = new(download.Checksum)
		checksum.Value, err = hex.DecodeString(checksumValue)
		if err != nil {
			return err
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
			return download.ErrChecksumAlgoNotSupported
		}
	}

	// Download File
	var decompressorDst string
	if decompressor != nil {
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

	parsedurl, err := url.Parse(s.Url)
	if err != nil {
		return err
	}

	getter := &download.HttpGetter{Client: dl.GetHttpClient()}
	if err := getter.Get(ctx, parsedurl, dst); err != nil {
		return err
	}
	if checksum != nil {
		if err := checksum.Checksum(dst); err != nil {
			return err
		}
	}
	if decompressor != nil {
		if err := decompressor.Decompress(ctx, dst, decompressorDst); err != nil {
			return err
		}
	}
	return nil
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
