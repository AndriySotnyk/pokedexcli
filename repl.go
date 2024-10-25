package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AndriySotnyk/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL     *string
	previousURL *string
	client      pokeapi.Client
}

const cliName = "Pokedex"

func startRepl(cfg *config) {
	sc := bufio.NewScanner(os.Stdin)
	commandMap := returnCommandMap()

	for {
		printPrompt()
		sc.Scan()
		text := cleanInput(sc.Text())
		if command, exists := commandMap[text]; exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			printUnknown(text)
			continue
		}
	}
}

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func printUnknown(text string) {
	fmt.Printf("\"%s\": command not found\n", text)
}

func returnCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			callback:    commandClear,
		},
		"map": {
			name:        "map",
			description: "Show next locations",
			callback:    mapF,
		},
		"mapb": {
			name:        "map",
			description: "Show previuous locations",
			callback:    mapB,
		},
	}
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
