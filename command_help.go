package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to %v! These are commands available: \n", cliName)

	commandMap := returnCommandMap()

	for k, v := range commandMap {
		fmt.Printf("%s - %v\n", k, v.description)
	}

	return nil
}
