#Charmap

Character encodings in Go language. 
Supports a number of 8bit encodings. 
Provides Encode and Decode functions to convert a string from and to UTF-8 respectively.


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
Converts a string from UTF-8 to the specified encoding. Returns converted string. 

If the input string contains illegal characters for the specified encoding,
these characters will be replaced with a substitute character ('?') and
ErrInvalidCodepoint will be returned in error value.

If the specified encoding is unknown, it will return the input string and ErrUnknownEncoding

    func Decode(data string, encoding string) (string, error)
Converts a string from the specified encoding to UTF-8.  Returns converted string. 

If the input string contains illegal characters for the specified encoding,
these characters will be replaced with a substitute character (utf8.RuneError) and
ErrInvalidCodepoint will be returned in error value.

If the specified encoding is unknown, it will return the input string and ErrUnknownEncoding

    func List() []string
Returns names of all supported encodings as a slice of strings
