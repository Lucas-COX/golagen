package main

import (
	"log"
	"os"

	"Lucas-COX/golagen/internal"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var err error = nil
	var config *internal.Config = nil

	if len(os.Args) > 1 {
		config, err = internal.ReadConfigFile(&os.Args[1])
	} else {
		config, err = internal.ReadConfigFile(nil)
	}
	if err != nil {
		log.Fatalf("invalid configuration file: %s.\n", err.Error())
	}

	spew.Dump(config)
}
