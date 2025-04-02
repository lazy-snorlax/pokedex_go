package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lazy-snorlax/pokedex/internal/pokeapi"
)

func commandWalk(cfg *config, args ...string) error {
	fmt.Println("Walking along the route...")
	time.Sleep(5 * time.Second)

	avail_pokemon := make([]string, 0)
	avail_pokemon = append(avail_pokemon, "bulbasaur", "charmander", "squirtle")

	found := avail_pokemon[rand.Intn(len(avail_pokemon))]
	foundPkmn, err := cfg.pokeapiClient.GetPokemon(found)
	if err != nil {
		return err
	}

	pkmnBattle(cfg, foundPkmn)
	return nil
}

func pkmnBattle(cfg *config, wildPkmn pokeapi.Pokemon) error {
	partyPkmn := cfg.partyPokemon["pikachu"]
	generateImg(wildPkmn.Sprites.FrontDefault)
	generateImg(partyPkmn.Sprites.BackDefault)
	fmt.Printf("A wild %s appeared!\n", wildPkmn.Name)
	fmt.Printf("Sent out %s!\n", partyPkmn.Name)

	fmt.Printf("What should %s do?\n", partyPkmn.Name)
	return nil
}
