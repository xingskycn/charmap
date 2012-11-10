package charmap

import (
	"strings"
	"unicode/utf8"
)

var aliasesMap = make(map[string]string)
var codecsMap = make(map[string]codec)

type charmap struct {
	Name    string
	Aliases []string
	Codec   codec
}

type codec interface {
	Encode(data string) (string, error)
	Decode(data string) (string, error)
}

func register(c charmap) {
	codecsMap[c.Name] = c.Codec
	for _, alias := range c.Aliases {
		aliasesMap[alias] = c.Name
	}
}

type EncodingNotSupportedError string

func (e EncodingNotSupportedError) Error() string {
	return "encoding not supported: " + string(e)
}

type EncodeError string

func (e EncodeError) Error() string {
	return string(e)
}

type DecodeError string

func (e DecodeError) Error() string {
	return string(e)
}

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
	encoding = strings.ToLower(encoding)
	encoding = strings.Replace(encoding, "-", "_", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.Encode(data)
		return result, err
	}

	return data, EncodingNotSupportedError(encoding)
}

// Decode decodes the string data from encoding into utf8
// If encoding is not supported it returns unchanged string and EncodingNotSupportedError
func Decode(data string, encoding string) (string, error) {
	encoding = strings.ToLower(encoding)
	encoding = strings.Replace(encoding, "-", "_", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.Decode(data)
		return result, err
	}

	return data, EncodingNotSupportedError(encoding)
}

// some functions to simplify 8bit codecs definition 

func reverseByteRuneMap(m map[byte]rune) (newmap map[rune]byte) {
	newmap = make(map[rune]byte)
	for k, v := range m {
		newmap[v] = k
	}
	return
}

func mapBytesToRunes(cm map[byte]rune, data string) (string, error) {
	size := len(data)
	result := make([]byte, 0, size)
	buf := make([]byte, 6)
	var n int
	var err error

	for i := 0; i < size; i++ {
		c := data[i]
		if r, ok := cm[c]; ok {
			n = utf8.EncodeRune(buf, r)
			result = append(result, buf[:n]...)
		} else {
			err = DecodeError("cannot decode one or more bytes")
			n = utf8.EncodeRune(buf, utf8.RuneError)
			result = append(result, buf[:n]...)
		}
	}

	return string(result), err
}

func mapRunesToBytes(cm map[rune]byte, data string) (string, error) {
	size := len(data)
	result := make([]byte, 0, size)
	var err error

	for _, r := range data {
		if c, ok := cm[r]; ok {
			result = append(result, c)
		} else {
			err = EncodeError("cannot encode one or more runes")
			result = append(result, '?')
		}
	}

	return string(result), err
}
