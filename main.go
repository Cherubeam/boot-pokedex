package main

import (
	"time"

	"github.com/Cherubeam/boot-pokedex/internal/pokeapi"
)

var client = pokeapi.NewClient(5 * time.Second)

func main() {
	cfg := &config{
		commands:  getCommands(),
		apiClient: client,
		pokedex:   make(map[string]pokeapi.PokemonResponse),
	}
	startRepl(cfg)
}
