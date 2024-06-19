package install

import (
	"strings"

	"github.com/Lucino772/envelop/internal/utils"
)

type HttpProcessor struct{}

func (p *HttpProcessor) GetExportVars(s Source) any {
	return struct{ Destination string }{
		Destination: s.Destination,
	}
}

func (p *HttpProcessor) IterTasks(s Source, yield func(*SourceProcessorTask) bool) {
	getter := &HttpGetter{Url: s.Url}
	var decompressor Decompressor

	decompressors := map[string]Decompressor{
		"zip":    &ZipDecompressor{},
		"tar":    &TarDecompressor{},
		"tar:gz": &TarGzipDecompressor{},
	}

	if s.Url.Query().Has("checksum") {
		results := strings.Split(s.Url.Query().Get("checksum"), ":")
		s.Url.Query().Del("checksum")
		if len(results) == 2 {
			getter.Hasher = utils.NewHash(results[0])
			getter.ExpectedHash = results[1]
		}
	}

	if s.Url.Query().Has("format") {
		format := s.Url.Query().Get("format")
		s.Url.Query().Del("format")
		decompressor = decompressors[format]
	}

	task := &SourceProcessorTask{
		Path:         s.Destination,
		Getter:       getter,
		Decompressor: decompressor,
	}
	if !yield(task) {
		return
	}
}
