package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"Lucas-COX/golagen/internal"
)

var configFile string
var verbose bool

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	flag.StringVar(&configFile, "config", "./configs/golagen.yaml", "Path to the configuration file")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	var err error = nil
	var config *internal.Config = nil

	log.Info(internal.BuildInfo())
	config, err = internal.ReadConfigFile(configFile)
	if err != nil {
		log.WithField("config", configFile).Fatalf("invalid configuration file: %s.\n", err.Error())
	}

	if err = internal.CheckPrerequisities(); err != nil {
		log.Fatalf("unmet prerequisities: %s.\n", err.Error())
	}

	if err = Generate(*config); err != nil {
		log.Fatalf("an error occured while generating files: %s.\n", err.Error())
	}
}
