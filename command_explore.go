package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you need to provide location name")
	}
	nameOrId := args[0]
	location, err := cfg.client.GetLocation(nameOrId)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", nameOrId)
	fmt.Println("Pokemons found:")
	for _, pokemonEncounters := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounters.Pokemon.Name)
	}
	return nil
}
