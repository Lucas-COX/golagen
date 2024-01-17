package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	var err error = nil

	if len(os.Args) > 1 {
		err = ReadConfigFile(&os.Args[1])
	} else {
		err = ReadConfigFile(nil)
	}
	if err != nil {
		log.Fatalf("invalid configuration file: %s.\n", err.Error())
	}

	err = CheckConfigFile()
	if err != nil {
		log.Fatalf("invalid configuration file: %s.\n", err.Error())
	}

	fmt.Printf("%v\n", viper.AllKeys())
}
