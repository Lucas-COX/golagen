package utils

import "os"

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
