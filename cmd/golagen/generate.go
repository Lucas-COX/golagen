package main

import (
	"Lucas-COX/golagen/cmd/golagen/generate"
	"Lucas-COX/golagen/internal"
	"os"
	"os/exec"
	"path"
)

func Generate(config internal.Config) error {
	err := exec.Command("mkdir", "-p", internal.GetCachePath()).Run()
	if err != nil {
		return err
	}
	// TODO download template files from this repo
	if _, err = os.Stat(path.Join(internal.GetCachePath(), "templates")); os.IsNotExist(err) {
		err = generate.FetchTemplates()
	}
	if err != nil {
		return err
	}
	// TODO generate source files for each entry
	if err = generate.GenerateSources(config.Entries, config.Project); err != nil {
		return err
	}

	// TODO generate global events dir
	// TODO generate SAM samconfig.toml from sam config
	// TODO generate SAM template.yml from entries
	// TODO generate global Makefile
	return nil
}
