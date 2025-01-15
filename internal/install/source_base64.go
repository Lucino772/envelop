package install

import (
	"context"
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
)

type Base64Source struct {
	Content string `mapstructure:"content,omitempty"`
}

func (s *Base64Source) GetDownloaderOptions() []DownloaderOptFunc { return nil }

func (s *Base64Source) Download(ctx context.Context, dl *Downloader, dst string) error {
	content, err := base64.URLEncoding.DecodeString(s.Content)
	if err != nil {
		return err
	}
	_, err = os.Stat(dst)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			return err
		}
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		if dstFile != nil {
			dstFile.Close()
		}
		return err
	}
	defer dstFile.Close()
	if _, err := dstFile.Write(content); err != nil {
		return err
	}
	return nil
}
