package download

import (
	"archive/zip"
	"context"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

type ZipDecompressor struct{}

func (dc *ZipDecompressor) Decompress(ctx context.Context, src string, dst string) error {
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

	rd, err := zip.OpenReader(src)
	if err != nil {
		if rd != nil {
			rd.Close()
		}
		return err
	}

	for _, file := range rd.File {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := extractZipFile(file, dst); err != nil {
				return err
			}
		}
	}
	return nil
}

func extractZipFile(file *zip.File, dst string) error {
	dstPath := path.Join(dst, file.Name)
	if !strings.HasPrefix(dstPath, path.Clean(dst)+string(os.PathSeparator)) {
		return errors.New("invalid file path")
	}

	if file.FileInfo().IsDir() {
		return os.MkdirAll(dstPath, os.ModePerm)
	}

	if err := os.MkdirAll(path.Dir(dstPath), os.ModePerm); err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		if dstFile != nil {
			dstFile.Close()
		}
		return err
	}
	defer dstFile.Close()
	srcFile, err := file.Open()
	if err != nil {
		if srcFile != nil {
			srcFile.Close()
		}
		return err
	}
	defer srcFile.Close()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}
