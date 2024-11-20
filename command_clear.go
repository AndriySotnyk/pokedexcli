package main

import (
	"os"
	"os/exec"
)

func commandClear(cfg *config, args ...string) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
