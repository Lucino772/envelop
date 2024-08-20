package wrapper

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

type HttpHook struct {
	Url string `mapstructure:"url,omitempty"`
}

func NewHttpHook(options map[string]any) *HttpHook {
	var hook HttpHook
	if err := mapstructure.Decode(options, &hook); err != nil {
		return nil
	}
	// TODO: Check url is valid
	return &hook
}

func (hook *HttpHook) Execute(parent context.Context, data []byte) error {
	// TODO: What about security/authentication
	ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", hook.Url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: Do we except a response ? If so, what's the shape ?
	return nil
}
