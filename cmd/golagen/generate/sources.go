package generate

import (
	"Lucas-COX/golagen/internal"
	"Lucas-COX/golagen/pkg/template"
	"Lucas-COX/golagen/pkg/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func copyWithReplacement(file string, dest string, project internal.Project) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	content, err = template.ApplyReplacements(content, map[string]string{
		"module_name": fmt.Sprintf("%s/%s", project.Author, project.Name),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, content, 0644)
	return err
}

func generateSource(entry internal.Entry, project internal.Project, path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		templatePath := filepath.Join(path, f.Name())
		if f.IsDir() {
			err = utils.ExecuteInDirectory(func() error {
				return generateSource(entry, project, templatePath)
			}, f.Name())
			if err != nil {
				return err
			}
		} else {
			err = copyWithReplacement(templatePath, f.Name(), project)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateSources(entries *[]internal.Entry, project internal.Project) error {
	var wg sync.WaitGroup
	if entries == nil || len(*entries) == 0 {
		return nil
	}

	templateDir := filepath.Join(internal.GetCachePath(), "templates", "go1.x")
	for _, entry := range *entries {
		log.Printf("Generating source files for %s...\n", entry.Name)
		wg.Add(1)
		go func(e internal.Entry) {
			defer wg.Done()
			err := utils.ExecuteInDirectory(func() error {
				return generateSource(e, project, templateDir)
			}, e.Name)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}(entry)
	}
	wg.Wait()
	return nil
}
