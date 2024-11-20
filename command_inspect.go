package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you need to provide Pokemon name")
	}
	name := args[0]
	if pokemon, ok := cfg.CaughtPokemon[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("	-%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("	- %v\n", pokeType.Type.Name)
		}
	} else {
		return fmt.Errorf("you have not caught this pokemon")
	}
	return nil
}
