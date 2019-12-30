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
