package main

import (
	"time"

	"github.com/AndriySotnyk/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
		client:        pokeClient,
	}
	startRepl(cfg)
}
