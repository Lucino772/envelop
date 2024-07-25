package steamweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	Root string
}

func NewClient() *Client {
	return &Client{Root: "api.steampowered.com"}
}

func (client *Client) Url(iface string, method string, version int, query url.Values) url.URL {
	return url.URL{
		Scheme:   "https",
		Host:     client.Root,
		Path:     fmt.Sprintf("%s/%s/v%d/", iface, method, version),
		RawQuery: query.Encode(),
	}
}

func (client *Client) CallJson(u url.URL, body any) error {
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, body)
}

func (client *Client) CallProto(u url.URL, body proto.Message) error {
	query := u.Query()
	query.Add("format", "protobuf_raw")
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return proto.Unmarshal(data, body)
}
