package main

import (
	"Lucas-COX/golagen/internal"
	"Lucas-COX/golagen/pkg/utils"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
)

var cache_path = ".golagen"
var template_repo_url = "git@github.com:Lucas-COX/golagen-templates.git"

func generateSource(entry internal.Entry, wg *sync.WaitGroup) error {
	var cmds []*exec.Cmd

	if wg != nil {
		defer wg.Done()
	}
	cmds = append(cmds, utils.CreateCommand(
		exec.Command("mkdir", "-p", entry.Name),
		&utils.CommandConfig{Cwd: "."},
	))
	return utils.RunCommandList(cmds)
}

func fetchTemplates() error {
	var cmds []*exec.Cmd

	cmds = append(cmds, utils.CreateCommand(
		exec.Command("git", "clone", "--depth", "1", "--no-checkout", template_repo_url, "templates"),
		&utils.CommandConfig{Cwd: cache_path},
	))
	cmds = append(cmds, utils.CreateCommand(
		exec.Command("git", "sparse-checkout", "set", "go1.x", "sam"),
		&utils.CommandConfig{Cwd: path.Join(cache_path, "templates")},
	))
	cmds = append(cmds, utils.CreateCommand(
		exec.Command("git", "checkout"),
		&utils.CommandConfig{Cwd: path.Join(cache_path, "templates")},
	))

	log.Println("Fetching aws templates...")
	return utils.RunCommandList(cmds)
}

func generateSources(entries *[]internal.Entry) error {
	var wg sync.WaitGroup
	if entries == nil {
		return nil
	}
	for _, entry := range *entries {
		log.Printf("Generating source files for %s...\n", entry.Name)
		wg.Add(1)
		go generateSource(entry, &wg)
	}
	wg.Wait()
	return nil
}

func Generate(config internal.Config) error {
	err := exec.Command("mkdir", "-p", cache_path).Run()
	if err != nil {
		return err
	}
	// TODO download template files from this repo
	if _, err = os.Stat(path.Join(cache_path, "templates")); os.IsNotExist(err) {
		err = fetchTemplates()
	}
	if err != nil {
		return err
	}
	// TODO generate source files for each entry
	if err = generateSources(config.Entries); err != nil {
		return err
	}

	// TODO generate global events dir
	// TODO generate SAM samconfig.toml from sam config
	// TODO generate SAM template.yml from entries
	// TODO generate global Makefile
	return nil
}
