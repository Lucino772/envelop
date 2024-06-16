package install

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/Lucino772/envelop/internal/utils"
)

type HttpGetter struct {
	Url  string
	Size uint32
	Hash struct {
		Algo  string
		Value string
	}
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

	resp, err := httpGetWithContext(ctx, g.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dstWriter := utils.NewChecksumFileWriter(dstFile, utils.NewHash(g.Hash.Algo))
	if _, err := io.CopyBuffer(
		dstWriter,
		utils.FullReader(resp.Body),
		make([]byte, 32*1024),
	); err != nil && err != io.EOF {
		return err
	}
	if err := dstWriter.Checksum(g.Hash.Value); err != nil {
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
