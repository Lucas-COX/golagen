package template

import (
	"Lucas-COX/golagen/internal"
	"Lucas-COX/golagen/pkg/utils"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

// Fetch the golagen template files that will be used for generating lambdas
func FetchGolagenTemplates() error {
	_, err := os.Stat(internal.GetCachePath())
	if err != nil {
		if err := os.MkdirAll(internal.GetCachePath(), 0755); err != nil {
			return err
		}
	}

	commands := [][]string{
		{"git", "clone", "--depth", "1", "-b", internal.GetVersion(), "--no-checkout", internal.GetTemplateRepoUrl(), "templates"},
		{"git", "-C", "templates", "sparse-checkout", "set", "go1.x", "sam"},
		{"git", "-C", "templates", "checkout", internal.GetVersion()},
	}

	_, err = os.Stat(path.Join(internal.GetCachePath(), "templates"))

	if err != nil {
		log.Info("Fetching aws templates")
		if err := utils.ExecuteCommands(commands, internal.GetCachePath()); err != nil {
			return err
		}
	}

	return nil
}
