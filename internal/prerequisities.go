package internal

import "os/exec"

func CheckPrerequisities() error {
	_, err := exec.LookPath("sam")

	if err != nil {
		return err
	}
	return nil
}
