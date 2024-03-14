package golagen

import (
	"Lucas-COX/golagen/internal"
	"Lucas-COX/golagen/pkg/template"
	"Lucas-COX/golagen/pkg/utils"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Entry struct {
	Name        string             `mapstructure:"name"`
	Methods     []string           `mapstructure:"methods"`
	Route       string             `mapstructure:"route"`
	Mods        *map[string]string `mapstructure:"mods,omitempty"`
	Environment *[]string          `mapstructure:"environment,omitempty"`
	Rules       *[]Rule            `mapstructure:"rules,omitempty"`
}

func (e *Entry) getLambdaTemplatesDir() string {
	return path.Join(internal.GetCachePath(), "templates", "go1.x")
}

func (e *Entry) getLambdaReplacements() map[string]string {
	project := viper.GetString("project")
	author := viper.GetString("author")

	return map[string]string{
		"module_name": fmt.Sprintf("%s/%s", author, project),
	}
}

func (e *Entry) copyLambdaTemplateFiles(file string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	destFile := utils.GenerateRelativePath(file, e.Name, "go1.x")
	if destFile == "" {
		return nil
	}
	if info.IsDir() {
		if err := os.MkdirAll(destFile, 0755); err != nil && !os.IsExist(err) {
			log.WithField("dest", destFile).Error("unable to create directory")
		}
		return nil
	}
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	content, err = template.ApplyReplacements(content, e.getLambdaReplacements(), destFile)
	if err != nil {
		return err
	}
	return os.WriteFile(destFile, content, 0755)
}

func (e *Entry) GenerateLambda() error {
	log.WithField("entry", e.Name).Info("generating source files")
	templatesDir := e.getLambdaTemplatesDir()
	// replacements := e.getLambdaReplacements()
	os.MkdirAll(e.Name, 0755)
	return filepath.Walk(templatesDir, e.copyLambdaTemplateFiles)
}

func GenerateLambdasAsync(entries []Entry) {
	var wg sync.WaitGroup

	for _, entry := range entries {
		wg.Add(1)
		go func(e Entry) {
			defer wg.Done()
			if err := e.GenerateLambda(); err != nil {
				log.WithField("entry", e.Name).Error(err.Error())
			}
		}(entry)
	}
	wg.Wait()
}
