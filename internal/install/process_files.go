package install

import (
	"path"
	"path/filepath"
)

type FilesProcessor struct {
	Files map[string]struct {
		Url  string `json:"url,omitempty"`
		Size uint32 `json:"size,omitempty"`
		Hash struct {
			Algo  string `json:"algo,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"hash,omitempty"`
	} `json:"files,omitempty"`
	Destination string                 `json:"destination,omitempty"`
	Exports     map[string]interface{} `json:"exports,omitempty"`
}

func (p *FilesProcessor) WithInstallDir(dir string) InstallProcessor {
	dst, err := filepath.Abs(filepath.Join(dir, p.Destination))
	if err != nil {
		return nil
	}

	return &FilesProcessor{
		Files:       p.Files,
		Destination: dst,
		Exports:     p.Exports,
	}
}

func (p *FilesProcessor) ParseExports() map[string]any {
	data := struct{ Destination string }{
		Destination: p.Destination,
	}
	return parseExports(p.Exports, data)
}

func (p *FilesProcessor) GetSize() uint32 {
	var size uint32 = 0
	for _, val := range p.Files {
		size += val.Size
	}
	return size
}

func (processor *FilesProcessor) IterTasks(yield func(*InstallTask) bool) {
	for filename, val := range processor.Files {
		task := &InstallTask{
			Path: path.Join(processor.Destination, filename),
			Getter: &HttpGetter{
				Url:  val.Url,
				Size: val.Size,
				Hash: struct {
					Algo  string
					Value string
				}{
					Algo:  val.Hash.Algo,
					Value: val.Hash.Value,
				},
			},
			Decompressor: nil,
		}
		if !yield(task) {
			return
		}
	}
}
