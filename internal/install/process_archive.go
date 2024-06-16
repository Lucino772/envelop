package install

import "path"

type ArchiveProcessor struct {
	Format string `json:"format,omitempty"`
	File   struct {
		Url  string `json:"url,omitempty"`
		Size uint32 `json:"size,omitempty"`
		Hash struct {
			Algo  string `json:"algo,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"hash,omitempty"`
	} `json:"file,omitempty"`
	Destination string `json:"destination,omitempty"`
}

func (p *ArchiveProcessor) WithInstallDir(dir string) InstallProcessor {
	return &ArchiveProcessor{
		Format:      p.Format,
		File:        p.File,
		Destination: path.Join(dir, p.Destination),
	}
}

func (p *ArchiveProcessor) GetSize() uint32 {
	return p.File.Size
}

func (p *ArchiveProcessor) IterTasks(yield func(*InstallTask) bool) {
	compressors := map[string]Decompressor{
		"zip": &ZipDecompressor{},
	}
	task := &InstallTask{
		Path: p.Destination,
		Getter: &HttpGetter{
			Url:  p.File.Url,
			Size: p.File.Size,
			Hash: struct {
				Algo  string
				Value string
			}{
				Algo:  p.File.Hash.Algo,
				Value: p.File.Hash.Value,
			},
		},
		Decompressor: compressors[p.Format],
	}
	if !yield(task) {
		return
	}
}
