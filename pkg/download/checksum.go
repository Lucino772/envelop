package download

import (
	"bytes"
	"errors"
	"hash"
	"io"
	"os"
)

var ErrChecksumMismatch = errors.New("checksum mismatch")

type Checksum struct {
	Value []byte
	Hash  hash.Hash
}

func (c *Checksum) Checksum(src string) error {
	file, err := os.Open(src)
	if err != nil {
		if file != nil {
			file.Close()
		}
		return err
	}
	defer file.Close()

	c.Hash.Reset()
	if _, err := io.Copy(c.Hash, file); err != nil {
		return err
	}

	if value := c.Hash.Sum(nil); !bytes.Equal(value, c.Value) {
		return ErrChecksumMismatch
	}
	return nil
}
