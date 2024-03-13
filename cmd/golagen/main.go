package main

import (
	"log"
	"os"

	"Lucas-COX/golagen/internal"
)

func main() {
	var err error = nil
	var config *internal.Config = nil

	log.Println(internal.BuildInfo())
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
