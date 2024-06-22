package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"os"
)

func GetSize(name string) (int64, error) {
	stat, err := os.Stat(name)
	if err != nil {
		return 0, nil
	}
	return stat.Size(), nil
}

func NewHash(name string) hash.Hash {
	switch name {
	case "md5":
		return md5.New()
	case "sha1":
		return sha1.New()
	case "sha256":
		return sha256.New()
	case "sha512":
		return sha512.New()
	}
	return nil
}

type readerWithCloseFn struct {
	io.Reader
	closeFn func() error
}

func WithCloseFn(r io.Reader, closeFn func() error) io.ReadCloser {
	return &readerWithCloseFn{r, closeFn}
}

func (r *readerWithCloseFn) Close() error {
	if r.closeFn != nil {
		return r.closeFn()
	}
	return nil
}

type fullReader struct {
	io.Reader
}

func NewFullReader(r io.Reader) io.Reader {
	return &fullReader{r}
}

func (r *fullReader) Read(buf []byte) (read int, err error) {
	read, err = io.ReadFull(r, buf)
	if err == io.ErrUnexpectedEOF {
		err = io.EOF
	}
	return read, err
}
