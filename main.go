package main

import (
	"time"

	"github.com/lazy-snorlax/pokedex/internal/pokeapi"
)

func main() {
	cacheInterval := 5 * time.Minute
	pokeClient := pokeapi.NewClient(5*time.Second, cacheInterval)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
