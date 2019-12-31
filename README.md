# go-pipe-parser

[![GoDoc](https://godoc.org/github.com/webermarci/go-pipe-parser?status.svg)](https://godoc.org/github.com/webermarci/go-pipe-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/webermarci/go-pipe-parser)](https://goreportcard.com/report/github.com/webermarci/go-pipe-parser)

All it does is just parser command strings to feed into `go-pipe`
https://github.com/b4b4r07/go-pipe

### Install
```
go get github.com/webermarci/go-pipe-parser
```

### Example
```go
package main

import (
	"log"
	pipeparser "webermarci/go-pipe-parser"
)

func main() {
	command := "ls -la | grep root"

	result, err := pipeparser.Run(command)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(result.String())
}
```
