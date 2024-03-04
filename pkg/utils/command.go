package utils

import "os/exec"

type CommandConfig struct {
	Cwd string
}

func CreateCommand(cmd *exec.Cmd, c *CommandConfig) *exec.Cmd {
	cmd.Dir = c.Cwd
	return cmd
}

func RunCommandList(cmds []*exec.Cmd) error {
	for _, cmd := range cmds {
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
