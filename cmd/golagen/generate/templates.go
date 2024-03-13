package generate

import (
	"Lucas-COX/golagen/internal"
	"Lucas-COX/golagen/pkg/utils"
	"log"
	"os"
)

func FetchTemplates() error {
	if err := os.MkdirAll(internal.GetCachePath(), 0755); err != nil {
		return err
	}

	commands := [][]string{
		{"git", "clone", "--depth", "1", "--no-checkout", internal.GetTemplateRepoUrl(), "templates"},
		{"git", "sparse-checkout", "set", "go1.x", "sam"},
		{"git", "checkout", internal.GetVersion()},
	}

	if err := utils.ExecuteCommands(commands, internal.GetCachePath()); err != nil {
		return err
	}

	log.Println("Fetching aws templates...")
	return nil
}
