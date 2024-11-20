package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you need to provide Pokemon name")
	}
	nameOrId := args[0]
	pokemon, err := cfg.client.GetPokemon(nameOrId)
	if err != nil {
		return err
	}
	if _, ok := cfg.CaughtPokemon[pokemon.Name]; ok {
		return fmt.Errorf("%s was already caught", pokemon.Name)
	}
	if rand.Intn(pokemon.BaseExperience) > 40 {
		cfg.CaughtPokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was catched\n", pokemon.Name)
	} else {
		fmt.Printf("%s was not caught\n", pokemon.Name)
	}
	return nil
}
