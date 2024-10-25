package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

const cliName = "Pokedex"

func startRepl() {
	sc := bufio.NewScanner(os.Stdin)
	commandMap := returnCommandMap()

	printPrompt()

	for sc.Scan() {
		text := cleanInput(sc.Text())
		if command, exists := commandMap[text]; exists {
			if command.name == "exit" {
				return
			}
			command.callback()
		} else {
			printUnknown(text)
		}
		printPrompt()
	}
}

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

func returnCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			callback:    commandClear,
		},
	}
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
