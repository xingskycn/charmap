package charmap

import (
	"testing"
	"unicode/utf8"
)

func TestEncode(t *testing.T) {
	pana_cp1251 := "\xC2\x20\xF7\xE0\xF9\xE0\xF5\x20\xFE\xE3\xE0\x20\xE6\xE8\xEB\x20\xE1\xFB\x20\xF6\xE8\xF2\xF0\xF3\xF1\x3F\x20\xC4\xE0\x2C\x20\xED\xEE\x20\xF4\xE0\xEB\xFC\xF8\xE8\xE2\xFB\xE9\x20\xFD\xEA\xE7\xE5\xEC\xEF\xEB\xFF\xF0\x21"
	pana_koi8r := "\xF7\x20\xDE\xC1\xDD\xC1\xC8\x20\xC0\xC7\xC1\x20\xD6\xC9\xCC\x20\xC2\xD9\x20\xC3\xC9\xD4\xD2\xD5\xD3\x3F\x20\xE4\xC1\x2C\x20\xCE\xCF\x20\xC6\xC1\xCC\xD8\xDB\xC9\xD7\xD9\xCA\x20\xDC\xCB\xDA\xC5\xCD\xD0\xCC\xD1\xD2\x21"
	pana_utf8 := "В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!"

	test_cp1251, err := Encode(pana_utf8, "cp1251")
	if err != nil {
		t.Error("encoding to cp1251: wrong error value")
	}
	if test_cp1251 != pana_cp1251 {
		t.Error("encoding to cp1251: wrong result")
	}

	test_koi8r, err := Encode(pana_utf8, "koi8-r")
	if err != nil {
		t.Error("encoding to koi8-r: wrong error value")
	}
	if test_koi8r != pana_koi8r {
		t.Error("encoding to koi8-r: wrong result")
	}

	test_alias, err := Encode(pana_utf8, "windows-1251")
	if err != nil {
		t.Error("encoding to alias windows-1251: wrong error value")
	}
	if test_alias != pana_cp1251 {
		t.Error("encoding to alias windows-1251: wrong result")
	}

	test_errUnEnc, err := Encode(pana_utf8, "wrong-encoding")
	if err != ErrUnknownEncoding {
		t.Error("encoding to wrong-encoding: wrong error value")
	}
	if test_errUnEnc != pana_utf8 {
		t.Error("encoding to wrong-encoding: wrong result")
	}

	test_illegal, err := Encode("AαZ", "cp1251")
	if err != ErrInvalidCodepoint {
		t.Error("encoding illegal codepoint: wrong error value")
	}
	if test_illegal != "A?Z" {
		t.Error("encoding illegal codepoint: wrong result")
	}

	test_empty, err := Encode("", "cp1251")
	if err != nil {
		t.Error("encoding empty string: wrong error value")
	}
	if test_empty != "" {
		t.Error("encoding empty string: wrong result")
	}
}

func TestDecode(t *testing.T) {
	pana_cp1251 := "\xC2\x20\xF7\xE0\xF9\xE0\xF5\x20\xFE\xE3\xE0\x20\xE6\xE8\xEB\x20\xE1\xFB\x20\xF6\xE8\xF2\xF0\xF3\xF1\x3F\x20\xC4\xE0\x2C\x20\xED\xEE\x20\xF4\xE0\xEB\xFC\xF8\xE8\xE2\xFB\xE9\x20\xFD\xEA\xE7\xE5\xEC\xEF\xEB\xFF\xF0\x21"
	pana_koi8r := "\xF7\x20\xDE\xC1\xDD\xC1\xC8\x20\xC0\xC7\xC1\x20\xD6\xC9\xCC\x20\xC2\xD9\x20\xC3\xC9\xD4\xD2\xD5\xD3\x3F\x20\xE4\xC1\x2C\x20\xCE\xCF\x20\xC6\xC1\xCC\xD8\xDB\xC9\xD7\xD9\xCA\x20\xDC\xCB\xDA\xC5\xCD\xD0\xCC\xD1\xD2\x21"
	pana_utf8 := "В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!"

	test_cp1251, err := Decode(pana_cp1251, "cp1251")
	if err != nil {
		t.Error("decoding from cp1251: wrong error value")
	}
	if test_cp1251 != pana_utf8 {
		t.Error("decoding from cp1251: wrong result")
	}

	test_koi8r, err := Decode(pana_koi8r, "koi8-r")
	if err != nil {
		t.Error("decoding from koi8-r: wrong error value")
	}
	if test_koi8r != pana_utf8 {
		t.Error("decoding from koi8-r: wrong result")
	}

	test_alias, err := Decode(pana_cp1251, "windows-1251")
	if err != nil {
		t.Error("decoding from alias windows-1251: wrong error value")
	}
	if test_alias != pana_utf8 {
		t.Error("decoding from alias windows-1251: wrong result")
	}

	test_errUnEnc, err := Decode(pana_utf8, "wrong-encoding")
	if err != ErrUnknownEncoding {
		t.Error("decoding from wrong-encoding: wrong error value")
	}
	if test_errUnEnc != pana_utf8 {
		t.Error("decoding from wrong-encoding: wrong result")
	}

	test_illegal, err := Decode("A\x98Z", "cp1251")
	if err != ErrInvalidCodepoint {
		t.Error("decoding illegal codepoint: wrong error value")
	}
	if test_illegal != "A"+string(utf8.RuneError)+"Z" {
		t.Error("decoding illegal codepoint: wrong result")
	}

	test_empty, err := Decode("", "cp1251")
	if err != nil {
		t.Error("encoding empty string: wrong error value")
	}
	if test_empty != "" {
		t.Error("encoding empty string: wrong result")
	}
}

func TestList(t *testing.T) {
	list := List()
	if len(list) == 0 {
		t.Error("list encoding: length === 0")
	}

	check := 0
	for _, enc := range list {
		if enc == "ISO-8859-5" || enc == "CP1251" || enc == "CP866" || enc == "KOI8-R" {
			check++
		}
	}
	if check != 4 {
		t.Error("list encoding: encodings not found in list")
	}
}
