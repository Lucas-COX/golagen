package utils

import (
	"os"
	"path"
	"strings"
)

func ExecuteInDirectory(fn func() error, dir string) error {
	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := os.Chdir(dir); err != nil {
		return err
	}

	fnerr := fn()

	if err := os.Chdir(originalDir); err != nil {
		return err
	}
	return fnerr
}

func GenerateRelativePath(base string, dest string, sep string) string {
	dir := path.Dir(base)
	index := strings.LastIndex(dir, sep)

	if index != -1 {
		result := base[index+len(sep):]
		return path.Join(".", dest, result)
	} else {
		return ""
	}
}
