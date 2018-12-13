# getld

A [golang][go] command line utility and library to extract [JSON-LD][jsonld] data embedded in [HTML][html].

## Installation

```sh
go get github.com/daetal-us/getld
```

## Usage
_Note: results are always returned as an [array][array]._

### Command Line

Retrieve [JSON-LD][jsonld] data embedded in [HTML][html] from a remote source:
```sh
getld http://apple.com
```

Retrieve [JSON-LD][jsonld] data embedded in [HTML][html] from multiple remote sources:

```sh
getld http://apple.com http://adobe.com
```
_Note: results from multiple remote sources are combined into a single [array][array]._

### Golang

```go
package main

import (
	"log"
	"encoding/json"

	"github.com/daetal-us/getld/extract"
)

func main() {
	extractFromHTML()
	extractFromURL()
}

func extractFromHTML() {
	html := `<html><body><script type="application/json+ld>{"@type":"organization","name":"example"}</script></body></html>"`
	results, err := extract.FromHTML(html)
	if err != nil {
		log.Fatal(err)
	}
	encoded, err := json.Marshal(results)
	if err != nil {
		log.fatal(err)
	}
	log.Println(string(encoded)) // [{"@type":"organization","name":"example"}]
}

func extractFromURL() {
	results, err := extract.FromURL("http://apple.com")
	if err != nil {
		log.Fatal(err)
	}
	encoded, err := json.Marshal(results)
	if err != nil {
		log.fatal(err)
	}
	log.Println(string(encoded)) // [...]
}
```

[go]:https://golang.org
[jsonld]:https://json-ld.org
[html]:https://www.w3.org/html
[array]:https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array