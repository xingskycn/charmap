package charmap

import (
	"strings"
	"unicode/utf8"
)

var aliasesMap = make(map[string]string)
var codecsMap = make(map[string]Codec)

type Codec interface {
	Encode(data string) string
	Decode(data string) string
}

type Charmap struct {
	Name    string
	Aliases []string
	Codec
}

func Register(c Charmap) {
	codecsMap[c.Name] = c.Codec
	for _, alias := range c.Aliases {
		aliasesMap[alias] = c.Name
	}
}

func List() []string {
	list := make([]string, 0)
	for name, _ := range codecsMap {
		list = append(list, name)
	}
	return list
}

func Encode(data string, encoding string) string {
	encoding = strings.ToLower(encoding)
	encoding = strings.Replace(encoding, "-", "_", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	if codec, ok := codecsMap[encoding]; ok {
		result := codec.Encode(data)
		return result
	}

	return ""
}

func Decode(data string, encoding string) string {
	encoding = strings.ToLower(encoding)
	encoding = strings.Replace(encoding, "-", "_", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	if codec, ok := codecsMap[encoding]; ok {
		result := codec.Decode(data)
		return result
	}

	return ""
}

// some functions to simplify 8bit codecs definition 

func reverseByteRuneMap(m map[byte]rune) (newmap map[rune]byte) {
	newmap = make(map[rune]byte)
	for k, v := range m {
		newmap[v] = k
	}
	return
}

func mapBytesToRunes_(cm map[byte]rune, data string) string {
	size := len(data)
	result := ""

	for i := 0; i < size; i++ {
		c := data[i]
		if r, ok := cm[c]; ok {
			result += string(r)
		}
	}

	return result
}

func mapBytesToRunes(cm map[byte]rune, data string) string {
	size := len(data)
	result := make([]byte, 0, size)
	buf := make([]byte, 6)
	var n int

	for i := 0; i < size; i++ {
		c := data[i]
		if r, ok := cm[c]; ok {
			n = utf8.EncodeRune(buf, r)
			result = append(result, buf[:n]...)
		}
	}

	return string(result)
}

func mapRunesToBytes(cm map[rune]byte, data string) string {
	size := len(data)
	result := make([]byte, 0, size)

	for _, r := range data {
		if c, ok := cm[r]; ok {
			result = append(result, c)
		}
	}

	return string(result)
}
