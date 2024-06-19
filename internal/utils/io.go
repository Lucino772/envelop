package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"io"
)

var ErrHashMismatch = errors.New("hash mismatch")

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

type fullReader struct {
	reader io.Reader
}

func FullReader(r io.Reader) io.Reader {
	return &fullReader{r}
}

func (rd *fullReader) Read(buf []byte) (read int, err error) {
	read, err = io.ReadFull(rd.reader, buf)
	if err == io.ErrUnexpectedEOF {
		err = io.EOF
	}
	return read, err
}
