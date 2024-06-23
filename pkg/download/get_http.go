package download

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/Lucino772/envelop/internal/utils"
)

type HttpGetter struct{}

func (g *HttpGetter) Get(ctx context.Context, u *url.URL, dst string) error {
	_, err := os.Stat(dst)
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

	resp, err := httpGetWithContext(ctx, u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := make([]byte, 32*1024)
	if _, err := io.CopyBuffer(
		dstFile,
		utils.NewFullReader(resp.Body),
		buf,
	); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func httpGetWithContext(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
