package install

import (
	"context"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/Lucino772/envelop/internal/utils"
)

var ErrHashMismatch = errors.New("hash mismatch")

type HttpGetter struct {
	Url          url.URL
	Hasher       hash.Hash
	ExpectedHash string
}

func (g *HttpGetter) Get(ctx context.Context, dstPath string) error {
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

	resp, err := httpGetWithContext(ctx, g.Url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body := utils.FullReader(resp.Body)
	if g.Hasher != nil {
		body = io.TeeReader(body, g.Hasher)
	}

	buf := make([]byte, 32*1024)
	if _, err := io.CopyBuffer(dstFile, body, buf); err != nil && err != io.EOF {
		return err
	}

	if g.Hasher != nil {
		if hex.EncodeToString(g.Hasher.Sum(nil)) != g.ExpectedHash {
			return ErrHashMismatch
		}
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
