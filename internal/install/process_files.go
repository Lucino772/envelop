package install

import (
	"path"
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
	Destination string `json:"destination,omitempty"`
}

func (p *FilesProcessor) WithInstallDir(dir string) InstallProcessor {
	return &FilesProcessor{
		Files:       p.Files,
		Destination: path.Join(dir, p.Destination),
	}
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
