package main

import (
	"github.com/marcussss1/simplevault/pkg/app"
	"log"
)

func main() {
	err := app.Run()
	log.Fatal(err)
	// TODO COPY [/bin/bash] DOCKERFILE
}
