package install

import (
	"context"
	"errors"
	"os"
	"path"
)

type DataGetter struct {
	Content []byte
}

func (g *DataGetter) Get(ctx context.Context, dstPath string) error {
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
	dstFile.Write(g.Content)
	return nil
}
