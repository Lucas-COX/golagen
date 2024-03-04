package main

import (
	"log"
	"os"

	"Lucas-COX/golagen/internal"
)

var (
	version   string
	buildTime string
)

func main() {
	var err error = nil
	var config *internal.Config = nil

	log.Printf("golagen version \"%s\", build %s", version, buildTime)
	if len(os.Args) > 1 {
		config, err = internal.ReadConfigFile(&os.Args[1])
	} else {
		config, err = internal.ReadConfigFile(nil)
	}
	if err != nil {
		log.Fatalf("invalid configuration file: %s.\n", err.Error())
	}

	if err = internal.CheckPrerequisities(); err != nil {
		log.Fatalf("unmet prerequisities: %s.\n", err.Error())
	}

	if err = Generate(*config); err != nil {
		log.Fatalf("an error occured while generating files: %s.\n", err.Error())
	}
}
