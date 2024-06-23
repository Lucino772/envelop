package download

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/url"
	"os"
	"strings"
)

var (
	ErrSchemeNotSupported        = errors.New("download not supported for given scheme")
	ErrArchiveFormatNotSupported = errors.New("archive format not supported")
	ErrInvalidChecksumFormat     = errors.New("invalid checksum format")
	ErrChecksumAlgoNotSupported  = errors.New("checksum algorithm not supported")
)

type Downloader struct {
	Src           string
	Dst           string
	Getters       map[string]Getter
	Decompressors map[string]Decompressor
}

func NewDownloader(src string, dst string) *Downloader {
	return &Downloader{
		Src:           src,
		Dst:           dst,
		Getters:       defaultGetters(),
		Decompressors: defaultDecompressors(),
	}
}

func (d *Downloader) Download(ctx context.Context) error {
	var dst string = d.Dst
	var decompressorDst string

	u, err := url.Parse(d.Src)
	if err != nil {
		return err
	}

	getter, ok := d.Getters[u.Scheme]
	if !ok {
		return ErrSchemeNotSupported
	}

	q := u.Query()

	var decompressor Decompressor
	if q.Has("archive") {
		archiveFormat := q.Get("archive")
		q.Del("archive")
		u.RawQuery = q.Encode()

		decompressor, ok = d.Decompressors[archiveFormat]
		if !ok {
			return ErrArchiveFormatNotSupported
		}

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

	var checksum *Checksum
	if q.Has("checksum") {
		values := strings.SplitN(q.Get("checksum"), ":", 2)
		q.Del("checksum")
		u.RawQuery = q.Encode()

		if len(values) != 2 {
			return ErrInvalidChecksumFormat
		}
		checksumAlgo, checksumValue := values[0], values[1]

		checksum = new(Checksum)
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
			return ErrChecksumAlgoNotSupported
		}
	}

	if err := getter.Get(ctx, u, dst); err != nil {
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
