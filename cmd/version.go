/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"Lucas-COX/golagen/internal"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version",
	Long: `This command displays the current version of golagen and
its build information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.BuildInfo())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
