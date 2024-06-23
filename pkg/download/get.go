package download

import (
	"context"
	"net/url"
)

type Getter interface {
	Get(ctx context.Context, u *url.URL, dst string) error
}

func defaultGetters() map[string]Getter {
	return map[string]Getter{
		"http":  &HttpGetter{},
		"https": &HttpGetter{},
		"data":  &DataGetter{},
	}
}
