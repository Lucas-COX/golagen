package utils

import "os/exec"

type CommandConfig struct {
	Cwd string
}

func ExecuteCommands(commands [][]string, dirs ...string) error {
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}

	for _, cmd := range commands {
		c := exec.Command(cmd[0], cmd[1:]...)
		c.Dir = dirs[0]
		if err := c.Run(); err != nil {
			return err
		}
	}
	return nil
}
