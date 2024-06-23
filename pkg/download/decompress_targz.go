package download

import (
	"compress/gzip"
	"context"
	"errors"
	"os"
)

type TarGzipDecompressor struct{}

func (dc *TarGzipDecompressor) Decompress(ctx context.Context, src string, dst string) error {
	info, err := os.Stat(dst)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			return err
		}
	} else if !info.IsDir() {
		return errors.New("destination path is not a directory")
	}
	r, err := os.Open(src)
	if err != nil {
		if r != nil {
			r.Close()
		}
		return err
	}
	defer r.Close()
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	return untar(ctx, gzr, dst)
}
