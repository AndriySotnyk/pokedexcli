package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("this command doesn't accept arguments")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf("	- %s\n", pokemon.Name)
	}
	return nil
}
