package download

import (
	"context"
	"encoding/base64"
	"errors"
	"net/url"
	"os"
	"path"
	"strings"
)

type DataGetter struct{}

func (g *DataGetter) Get(ctx context.Context, u *url.URL, dst string) error {
	url := strings.Split(u.RequestURI(), ";")
	// TODO: Handle different types of media types
	contentBase64 := strings.Split(url[len(url)-1], ",")[1]
	content, err := base64.URLEncoding.DecodeString(contentBase64)
	if err != nil {
		return err
	}
	_, err = os.Stat(dst)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(path.Dir(dst), os.ModePerm); err != nil {
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
