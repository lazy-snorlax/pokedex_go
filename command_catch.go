package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if chance > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	pokemonImg := pokemon.Sprites.FrontDefault
	// fmt.Println("Image: ", pokemonImg)
	generateImg(pokemonImg)
	fmt.Println("You may now inspect it with the inspect command.")

	if len(cfg.partyPokemon) == 6 {
		cfg.storedPokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was sent to the PC.\n", pokemon.Name)
	} else {
		cfg.partyPokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was added to the party.\n", pokemon.Name)
	}
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}

func generateImg(imgUrl string) {
	flags := aic_package.DefaultFlags()
	flags.Colored = true
	flags.Dimensions = []int{50, 20}

	// Conversion for an image
	asciiArt, err := aic_package.Convert(imgUrl, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
}
