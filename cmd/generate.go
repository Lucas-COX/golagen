/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Lucas-COX/golagen/pkg/golagen"
	"Lucas-COX/golagen/pkg/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the monorepo files from configuration",
	Long: `Generate the lambda sample folders and the SAM template files if
they don't exist, then update those files depending on the configuration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var entries []golagen.Entry

		if err := template.FetchGolagenTemplates(); err != nil {
			return err
		}
		if err := viper.UnmarshalKey("entries", &entries); err != nil {
			return err
		}

		golagen.GenerateLambdasAsync(entries)

		// TODO generate source files for each entry
		// if err = generate.GenerateSources(golagen.Entries, config.Project); err != nil {
		// 	return err
		// }

		// TODO generate global events dir
		// TODO generate SAM samconfig.toml from sam config
		// TODO generate SAM template.yml from entries
		// TODO generate global Makefile
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
