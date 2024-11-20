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
	callback    func(*config, ...string) error
}

type config struct {
	nextURL       *string
	previousURL   *string
	client        pokeapi.Client
	CaughtPokemon map[string]pokeapi.Pokemon
}

const cliName = "Pokedex"

func startRepl(cfg *config) {
	sc := bufio.NewScanner(os.Stdin)
	commandMap := returnCommandMap()

	for {
		printPrompt()
		sc.Scan()
		args := strings.Split(cleanInput(sc.Text()), " ")
		if command, exists := commandMap[args[0]]; exists {
			if err := sc.Err(); err != nil {
				fmt.Println("Error, reading input", err)
				commandMap["exit"].callback(cfg, args[1])
			}
			var err error
			switch {
			case len(args) == 2:
				err = command.callback(cfg, args[1])
			default:
				err = command.callback(cfg)
			}
			if err != nil {
				fmt.Println(err)
			}
		} else {
			printUnknown(args[0])
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
			callback:    commandMapF,
		},
		"mapb": {
			name:        "map",
			description: "Show previuous locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Show found",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all pokemons in pokedex",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
