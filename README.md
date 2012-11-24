#Charmap

Character encodings in Go.



###Installation
    go get github.com/disintegration/charmap
    
###Code example

```go
package main

import (
    "github.com/disintegration/charmap"
	"io/ioutil"
)

func main() {
	// read text file and convert string from CP1251 encoding to UTF-8
	data, _ := ioutil.ReadFile("text-cp1251.txt")
	str, _ := charmap.Decode(string(data), "cp1251")

	// do something with str here

	// convert string from UTF-8 to KOI8-R encoding and save to file
	strkoi8r, _ := charmap.Encode(str, "koi8-r")
	ioutil.WriteFile("text-koi8-r.txt", []byte(strkoi8r), 0666)
}

```

###Overview

    func Encode(data string, encoding string) (string, error)
Converts the string from UTF-8 to the specified encoding. 

If the string contains illegal characters for this encoding,
these characters are replaced with a substitute character ('?') and
ErrInvalidCodepoint is returned in second return value.

In case of unknown encoding, returned string remains unchanged and
ErrUnknownEncoding is returned in second return value

    func Decode(data string, encoding string) (string, error)
Converts string from the specified encoding to UTF-8 

If the string contains illegal characters for this encoding,
these characters are replaced with a substitute character (utf8.RuneError) and
ErrInvalidCodepoint is returned in second return value

In case of unknown encoding, returned string remains unchanged and
ErrUnknownEncoding is returned in second return value

    func List() []string
Returns names of all supported encodings as a slice of strings
