package download

import "context"

type Decompressor interface {
	Decompress(ctx context.Context, src string, dst string) error
}

func defaultDecompressors() map[string]Decompressor {
	return map[string]Decompressor{
		"zip":    &ZipDecompressor{},
		"tar":    &TarDecompressor{},
		"tar:gz": &TarGzipDecompressor{},
	}
}
