package main

import (
	"log"

	"github.com/Lucino772/envelop/internal/cli"
)

func main() {
	if err := cli.RootCommand().Execute(); err != nil {
		log.Fatal("An error occured !")
	}
}
