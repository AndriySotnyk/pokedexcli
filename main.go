package main

import (
	"time"

	"github.com/AndriySotnyk/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second * 5)
	cfg := &config{
		client: pokeClient,
	}
	startRepl(cfg)
}
