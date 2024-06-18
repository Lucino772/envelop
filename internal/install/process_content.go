package install

import (
	"path"
	"path/filepath"
)

type ContentProcessor struct {
	Files       map[string]string      `json:"files,omitempty"`
	Destination string                 `json:"destination,omitempty"`
	Exports     map[string]interface{} `json:"exports,omitempty"`
}

func (p *ContentProcessor) WithInstallDir(dir string) InstallProcessor {
	dst, err := filepath.Abs(filepath.Join(dir, p.Destination))
	if err != nil {
		return nil
	}

	return &ContentProcessor{
		Files:       p.Files,
		Destination: dst,
		Exports:     p.Exports,
	}
}

func (p *ContentProcessor) ParseExports() map[string]any {
	data := struct{ Destination string }{
		Destination: p.Destination,
	}
	return parseExports(p.Exports, data)
}

func (p *ContentProcessor) GetSize() uint32 {
	var size uint32 = 0
	for _, content := range p.Files {
		size += uint32(len(content))
	}
	return size
}

func (p *ContentProcessor) IterTasks(yield func(*InstallTask) bool) {
	for filename, content := range p.Files {
		task := &InstallTask{
			Path: path.Join(p.Destination, filename),
			Getter: &ContentGetter{
				Content: content,
				Size:    uint32(len(content)),
			},
			Decompressor: nil,
		}
		if !yield(task) {
			return
		}
	}
}
