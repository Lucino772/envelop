package install

import "path"

type ContentProcessor struct {
	Files       map[string]string `json:"files,omitempty"`
	Destination string            `json:"destination,omitempty"`
}

func (p *ContentProcessor) WithInstallDir(dir string) InstallProcessor {
	return &ContentProcessor{
		Files:       p.Files,
		Destination: path.Join(dir, p.Destination),
	}
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
