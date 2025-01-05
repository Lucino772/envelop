package install

import (
	"context"
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	"github.com/alitto/pond/v2"
)

type Base64Source struct {
	Type        string         `mapstructure:"type,omitempty"`
	Destination string         `mapstructure:"destination,omitempty"`
	Exports     map[string]any `mapstructure:"exports,omitempty"`
	Content     string         `mapstructure:"url,omitempty"`
}

func (s *Base64Source) GetDownloaderOptions() []DownloaderOptFunc { return nil }

func (s *Base64Source) GetMetadata(ctx context.Context, dlCtx DownloadContext, dl *Downloader) (Metadata, error) {
	content, err := base64.URLEncoding.DecodeString(s.Content)
	if err != nil {
		return nil, err
	}
	return &Base64SourceMetadata{
		Content:     content,
		Destination: filepath.Join(dlCtx.InstallDir, s.Destination),
		Exports:     s.Exports,
		ContentLen:  int64(len(content)),
	}, nil
}

type Base64SourceMetadata struct {
	Content     []byte
	ContentLen  int64
	Destination string
	Exports     map[string]any
}

func (metadata *Base64SourceMetadata) GetExports() map[string]any {
	data := struct{ Destination string }{
		Destination: metadata.Destination,
	}
	return parseExports(metadata.Exports, data)
}

func (metadata *Base64SourceMetadata) Install(ctx context.Context, pool pond.Pool, dl *Downloader) (Waiter, error) {
	// TODO: Add decompression in case the data was compressed and then encode to base64
	// TODO: Add a checksum to validate the data

	group := pool.NewGroup()
	group.SubmitErr(func() error {
		_, err := os.Stat(metadata.Destination)
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(filepath.Dir(metadata.Destination), os.ModePerm); err != nil {
				return err
			}
		}
		dstFile, err := os.Create(metadata.Destination)
		if err != nil {
			if dstFile != nil {
				dstFile.Close()
			}
			return err
		}
		defer dstFile.Close()
		if _, err := dstFile.Write(metadata.Content); err != nil {
			return err
		}
		return nil
	})
	return group, nil
}
