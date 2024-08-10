package wrapper

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

type Hook interface {
	Execute(context.Context, []byte) error
}

func NewHook(typ string, options map[string]any) Hook {
	switch typ {
	case "http":
		return NewHttpHook(options)
	default:
		return nil
	}
}

// Http Web Hook
type httpHook struct {
	url string
}

func NewHttpHook(options map[string]any) *httpHook {
	url, ok := options["url"]
	if !ok {
		return nil
	}
	return &httpHook{url: url.(string)}
}

func (hook *httpHook) Execute(parent context.Context, data []byte) error {
	// TODO: What about security

	ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", hook.url, bytes.NewBuffer(data))
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
