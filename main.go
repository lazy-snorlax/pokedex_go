package main

import (
	"time"

	"github.com/lazy-snorlax/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		partyPokemon:  map[string]pokeapi.Pokemon{},
		storedPokemon: map[string]pokeapi.Pokemon{},
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
