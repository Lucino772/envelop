package steamvdf

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	TokenTypeString TokenType = iota
	TokenTypeChildStart
	TokenTypeChildEnd
)

type TokenType byte

type Token struct {
	Type  TokenType
	Value string
}

// Thanks to https://github.com/marshauf/keyvalues
func readText(data []byte) (kv KeyValue, err error) {

	data = bytes.TrimSpace(data)

	if len(data) == 0 {
		return KeyValue{}, nil
	}

	line := 1
	var tokens []*Token

	for i := 0; i < len(data); i++ {

		switch data[i] {
		case '"':

			j := i + 1

			var escaping bool

			for ; ; j++ {
				if j >= len(data) {
					return kv, errors.New("EOF")
				}

				if data[j] == '\\' {
					escaping = !escaping

					continue
				}

				if data[j] == '"' {
					if escaping {
						escaping = false

						continue
					}

					break
				}
			}

			str := string(data[i+1 : j])

			tokens = append(tokens, &Token{Type: TokenTypeString, Value: str})

			i = j

		case ' ', '\t':

			continue

		case '\n', '\r':

			line++
			continue

		case '{':

			tokens = append(tokens, &Token{Type: TokenTypeChildStart, Value: "{"})

		case '}':

			tokens = append(tokens, &Token{Type: TokenTypeChildEnd, Value: "}"})

		case byte(0):

			break

		case '/', '#':

			for ; i < len(data); i++ {

				if data[i] == '\n' || data[i] == '\r' {
					line++

					break
				}
			}

		default:

			i--

			if len(tokens) != 0 {
				return kv, fmt.Errorf("unhandled char \"%s\" at char %d in line %d\nlast token: %v", string(data[i]), i, line, tokens[len(tokens)-1])
			}

			return kv, fmt.Errorf("unhandled char \"%s\" at char %d in line %d", string(data[i]), i, line)
		}
	}

	root := KeyValue{}
	readObject(tokens, &root)

	return root.Children[0], nil // Return root or first root Children? Is it possible that multiple top level key, value pairs can exist?
}

func readObject(tokens []*Token, root *KeyValue) int {

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		case TokenTypeString:
			// Peek ahead
			switch tokens[i+1].Type {
			case TokenTypeString: // key, value (string)
				root.SetChild(KeyValue{
					Key:   tokens[i].Value,
					Value: tokens[i+1].Value})
				i++
			case TokenTypeChildStart: // key, value (object)
				child := KeyValue{Key: tokens[i].Value}
				read := readObject(tokens[i+2:], &child)
				root.SetChild(child)
				i += 2 + read
			}
		case TokenTypeChildEnd:
			return i
		default:
			panic(fmt.Errorf("unknown token type. %v", tokens[i]))
		}
	}
	return len(tokens)
}
