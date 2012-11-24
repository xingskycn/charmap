package charmap

import (
	"bytes"
	"errors"
	"strings"
	"unicode/utf8"
)

var aliasesMap = make(map[string]string)
var codecsMap = make(map[string]codec)

type codec interface {
	Encode(data string) (string, error)
	Decode(data string) (string, error)
}

func register(c codec, name string, aliases ...string) {
	codecsMap[name] = c
	for _, alias := range aliases {
		aliasesMap[alias] = name
	}
}

var ErrUnknownEncoding error = errors.New("encoding is not supported")
var ErrInvalidCodepoint error = errors.New("cannot convert one or more codepoints")

// List returns a list of all supported encodings as a slice of strings
func List() []string {
	list := make([]string, 0)
	for name, _ := range codecsMap {
		list = append(list, name)
	}
	return list
}

// Encode encodes the string data from utf8 to encoding
// If encoding is not supported it returns unchanged string and EncodingNotSupportedError
func Encode(data string, encoding string) (string, error) {
	encoding = strings.ToUpper(encoding)
	encoding = strings.Replace(encoding, "_", "-", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.Encode(data)
		return result, err
	}

	return data, ErrUnknownEncoding
}

// Decode decodes the string data from encoding into utf8
// If encoding is not supported it returns unchanged string and EncodingNotSupportedError
func Decode(data string, encoding string) (string, error) {
	encoding = strings.ToUpper(encoding)
	encoding = strings.Replace(encoding, "_", "-", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.Decode(data)
		return result, err
	}

	return data, ErrUnknownEncoding
}

// simple 8bit codecs definition support

func reverseByteRuneMap(m map[byte]rune) (newmap map[rune]byte) {
	newmap = make(map[rune]byte)
	for k, v := range m {
		newmap[v] = k
	}
	return
}

func mapBytesToRunes(cm map[byte]rune, data string) (result string, err error) {
	size := len(data)
	buf := bytes.NewBuffer(make([]byte, 0, size))

	for i := 0; i < size; i++ {
		c := data[i]
		if r, ok := cm[c]; ok {
			buf.WriteRune(r)
		} else {
			err = ErrInvalidCodepoint
			buf.WriteRune(utf8.RuneError)
		}
	}

	result = buf.String()
	return result, err
}

func mapRunesToBytes(cm map[rune]byte, data string) (result string, err error) {
	size := len(data)
	buf := bytes.NewBuffer(make([]byte, 0, size/2))

	for _, r := range data {
		if c, ok := cm[r]; ok {
			buf.WriteByte(c)
		} else {
			err = ErrInvalidCodepoint
			buf.WriteByte('?')
		}
	}

	result = buf.String()
	return result, err
}

type codecMap8Bit struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c *codecMap8Bit) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c *codecMap8Bit) Decode(s string) (string, error) {
	return mapBytesToRunes(c.DecodeMap, s)
}
