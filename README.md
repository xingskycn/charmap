#Charmap

Character encodings support in Go.



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
  // read text file and convert string from CP1251 encoding to UTF8
  data, _ := ioutil.ReadFile("sometext-win1251.txt")
  str, _ := charmap.Decode(string(data), "cp1251")
  
  // do something with the string here
  
  // convert string from UTF8 to KOI8R encoding and save to file
  strkoi8r, _ := charmap.Encode(str, "koi8r")
  ioutil.WriteFile("test1.txt", []byte(strkoi8r), 0666)
}

```


