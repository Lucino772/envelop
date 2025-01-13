package install

import (
	"context"
	"path/filepath"

	"github.com/alitto/pond/v2"
)

type FilesDepotManifest struct {
	Files []FilesDepotManifestFile `mapstructure:"files,omitempty"`
}

type FilesDepotManifestFile struct {
	Filename string `mapstructure:"filename,omitempty"`
	Source   Source `mapstructure:"source,omitempty"`
}

func (manifest *FilesDepotManifest) GetDownloaderOptions() []DownloaderOptFunc {
	var opts = make([]DownloaderOptFunc, 0, len(manifest.Files))
	for _, file := range manifest.Files {
		opts = append(opts, file.Source.GetDownloaderOptions()...)
	}
	return opts
}

func (manifest *FilesDepotManifest) GetMetadata(ctx context.Context, dl *Downloader, path string) (Metadata, error) {
	files := make([]FilesDepotManifestFile, 0, len(manifest.Files))
	for _, file := range manifest.Files {
		files = append(files, FilesDepotManifestFile{
			Filename: filepath.ToSlash(filepath.Join(path, file.Filename)),
			Source:   file.Source,
		})
	}

	return &FilesDepotManifestMetadata{Files: files}, nil
}

type FilesDepotManifestMetadata struct {
	Files []FilesDepotManifestFile
}

func (metadata *FilesDepotManifestMetadata) Install(ctx context.Context, pool pond.Pool, dl *Downloader) (Waiter, error) {
	rootGroup := pool.NewGroupContext(ctx)
	for _, file := range metadata.Files {
		rootGroup.SubmitErr(func() error {
			return file.Source.Download(ctx, dl, file.Filename)
		})
	}
	return rootGroup, nil
}
