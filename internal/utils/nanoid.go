package utils

import "crypto/rand"

const (
	nanoIdSize       = 21
	nanoIdBufferSize = nanoIdSize * nanoIdSize * 7
)

var nanoIdStandardAlphabet = [...]byte{
	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's',
	't', 'u', 'v', 'w', 'x',
	'y', 'z', '0', '1', '2',
	'3', '4', '5', '6', '7',
	'8', '9', '-', '_',
}

func NewNanoIDGenerator() (func() (string, error), error) {
	buf := make([]byte, nanoIdBufferSize)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	offset := 0

	var id [nanoIdSize]byte

	return func() (string, error) {
		// Refill if all the bytes have been used.
		if offset >= nanoIdBufferSize {
			if _, err := rand.Read(buf); err != nil {
				return "", err
			}

		}

		for i := 0; i < nanoIdSize; i++ {
			/*
				"It is incorrect to use bytes exceeding the alphabet size.
				The following mask reduces the random byte in the 0-255 value
				range to the 0-63 value range. Therefore, adding hacks such
				as empty string fallback or magic numbers is unneccessary because
				the bitmask trims bytes down to the alphabet size (64)."
			*/
			id[i] = nanoIdStandardAlphabet[buf[i+offset]&63]
		}

		offset += nanoIdSize
		return string(id[:]), nil
	}, nil
}
