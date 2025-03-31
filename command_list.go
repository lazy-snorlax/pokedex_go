package main

import (
	"errors"
	"fmt"
)

func commandList(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must choose either party or pc")
	}

	name := args[0]
	if name == "party" {
		fmt.Println("Your Party:")
		for _, pokemon := range cfg.partyPokemon {
			fmt.Printf(" - %s \n", pokemon.Name)
		}
	}
	if name == "pc" {
		fmt.Println("Your PC:")
		for _, pokemon := range cfg.storedPokemon {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}
	return nil
}
