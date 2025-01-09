package steamcdn

import (
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ClientOptFunc func(*cdnClient)
type cdnClient struct {
	httpClient *http.Client
}

func NewClient(opts ...ClientOptFunc) *cdnClient {
	client := &cdnClient{
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithHttpClient(httpClient *http.Client) ClientOptFunc {
	return func(c *cdnClient) {
		c.httpClient = httpClient
	}
}

func (client *cdnClient) DownloadManifest(server string, depotId uint32, manifestId uint64, manifestRequestCode uint64, depotKey []byte) (*DepotManifest, error) {
	var command string
	if manifestRequestCode > 0 {
		command = fmt.Sprintf("depot/%d/manifest/%d/%d/%d", depotId, manifestId, 5, manifestRequestCode)
	} else {
		command = fmt.Sprintf("depot/%d/manifest/%d/%d", depotId, manifestId, 5)
	}

	uri := client.buildUrl(server, command)
	manifestData, err := client.doCommand(uri)
	if err != nil {
		return nil, err
	}
	uncompressed, err := decompressZip(manifestData)
	if err != nil {
		return nil, err
	}
	return NewDepotManifest(uncompressed, depotKey)
}

func (client *cdnClient) DownloadDepotChunk(server string, depotId uint32, chunk ChunkData, depotKey []byte, cdnToken string) (*DepotChunk, error) {
	chunkId := hex.EncodeToString(chunk.ChunkId)
	uri := client.buildUrl(server, fmt.Sprintf("depot/%d/chunk/%s", depotId, chunkId))
	uri.RawQuery = cdnToken

	chunkData, err := client.doCommand(uri)
	if err != nil {
		return nil, err
	}
	return NewDepotChunk(chunk, chunkData, depotKey)
}

func (client *cdnClient) doCommand(uri url.URL) ([]byte, error) {
	request, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}
	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func (client *cdnClient) buildUrl(host string, command string) url.URL {
	return url.URL{
		Scheme: "https",
		Host:   host,
		Path:   command,
	}
}
