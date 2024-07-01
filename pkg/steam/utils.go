package steam

import "regexp"

func BitSet64_Set(b *uint64, v uint64, offset uint32, mask uint64) {
	*b = (*b & ^(mask << uint16(offset))) | ((v & mask) << uint16(offset))
}

func BitSet64_Get(b *uint64, offset uint32, mask uint64) uint64 {
	return (*b >> uint16(offset)) & mask
}

func processRegexGroups(regex *regexp.Regexp, matches []string) map[string]string {
	result := make(map[string]string)
	for ix, name := range regex.SubexpNames() {
		if ix != 0 && name != "" {
			result[name] = matches[ix]
		}
	}
	return result
}
