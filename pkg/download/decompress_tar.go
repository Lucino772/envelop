package download

import (
	"archive/tar"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type TarDecompressor struct{}

func (dc *TarDecompressor) Decompress(ctx context.Context, src string, dst string) error {
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

	return untar(ctx, r, dst)
}

func untar(ctx context.Context, r io.Reader, dst string) error {
	rd := tar.NewReader(r)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			header, err := rd.Next()
			switch {
			case err == io.EOF:
				return nil
			case err != nil:
				return err
			case header == nil:
				continue
			}
			if err := extractTarFile(rd, header, dst); err != nil {
				return err
			}
		}
	}
}

func extractTarFile(tr *tar.Reader, file *tar.Header, dst string) error {
	dstPath := filepath.Join(dst, file.Name)
	if !strings.HasPrefix(dstPath, filepath.Clean(dst)+string(os.PathSeparator)) {
		return errors.New("invalid file path")
	}
	if file.FileInfo().IsDir() {
		return os.MkdirAll(dstPath, os.ModePerm)
	}
	if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
		return err
	}
	dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(file.Mode))
	if err != nil {
		if dstFile != nil {
			dstFile.Close()
		}
		return err
	}
	defer dstFile.Close()
	if _, err := io.Copy(dstFile, tr); err != nil {
		return err
	}
	return nil
}
