package main

import (
	"log"

	"github.com/cuenobi/golang-clean/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
