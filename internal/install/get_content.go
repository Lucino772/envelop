package install

import (
	"context"
	"errors"
	"io"
	"os"
	"path"
)

type ContentGetter struct {
	Content string
	Size    uint32
}

func (g *ContentGetter) Get(ctx context.Context, dstPath string) error {
	_, err := os.Stat(dstPath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(path.Dir(dstPath), os.ModePerm); err != nil {
			return err
		}
	}
	dstFile, err := os.Create(dstPath)
	if err != nil {
		if dstFile != nil {
			dstFile.Close()
		}
		return err
	}
	defer dstFile.Close()
	io.WriteString(dstFile, g.Content)
	return nil
}
