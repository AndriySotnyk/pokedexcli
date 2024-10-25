package main

import (
	"os"
	"os/exec"
)

func commandClear(cfg *config) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
