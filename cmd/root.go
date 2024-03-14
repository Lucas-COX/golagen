/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Lucas-COX/golagen/pkg/golagen"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golagen",
	Short: "Generate a Go lambda monorepo with AWS SAM",
	Long: `A CLI tool that generates an AWS Lambda monorepo that uses the SAM cli
for deployment and resource creation.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLogging)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "configs/golagen.yaml", "config file (default is configs/golagen.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.SetDefault("entries", []golagen.Entry{})
	viper.SetDefault("environment", map[string]string{})

	if err := viper.ReadInConfig(); err != nil {
		log.WithField("config", cfgFile).Fatal("Unable to read in configuration")
	}
}

func initLogging() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
