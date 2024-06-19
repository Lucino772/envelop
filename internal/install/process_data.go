package install

import (
	"encoding/base64"
	"strings"
)

type DataProcessor struct{}

func (p *DataProcessor) GetExportVars(s Source) any {
	return struct{ Destination string }{
		Destination: s.Destination,
	}
}

func (p *DataProcessor) IterTasks(s Source, yield func(*SourceProcessorTask) bool) {
	url := strings.Split(s.Url.RequestURI(), ";")
	// TODO: Handle different types of media types

	content := strings.Split(url[len(url)-1], ",")[1]
	contentBytes, err := base64.URLEncoding.DecodeString(content)
	if err != nil {
		return
	}

	task := &SourceProcessorTask{
		Path: s.Destination,
		Getter: &DataGetter{
			Content: contentBytes,
		},
		Decompressor: nil,
	}
	if !yield(task) {
		return
	}
}
