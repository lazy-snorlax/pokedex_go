package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide one pokemon name")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
	} else {
		generateImg(pokemon.Sprites.FrontDefault)
		fmt.Printf("Name: %s \n", pokemon.Name)
		fmt.Printf("Height: %v \n", pokemon.Height)
		fmt.Printf("Weight: %v \n", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v \n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, pkType := range pokemon.Types {
			fmt.Printf("  -%v \n", pkType.Type.Name)
		}
	}

	return nil
}
