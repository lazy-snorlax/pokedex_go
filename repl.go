package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lazy-snorlax/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string

	partyPokemon  map[string]pokeapi.Pokemon
	storedPokemon map[string]pokeapi.Pokemon
	caughtPokemon map[string]pokeapi.Pokemon

	// bagItems map[string]pokeapi.Item
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokemon Terminal Edition")
	fmt.Println("===================================")
	fmt.Println("You are currently on Route 1.")
	fmt.Println("Use 'walk' to traverse the route. You may get a random encounter")
	fmt.Println("Use 'help' for a list of commands")
	starter, err := cfg.pokeapiClient.GetPokemon("pikachu")
	if err != nil {
		return
	}
	cfg.partyPokemon[starter.Name] = starter
	cfg.caughtPokemon[starter.Name] = starter

	for {
		fmt.Print("Player > ")
		reader.Scan()

		inputs := cleanInput(reader.Text())
		commandName := inputs[0]
		args := []string{}
		if len(inputs) > 1 {
			args = inputs[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"walk": {
			name:        "walk>",
			description: "Walk along a route",
			callback:    commandWalk,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inpsect a caught Pokemon",
			callback:    commandInspect,
		},
		"list": {
			name:        "list <party_or_pc>",
			description: "List party Pokemon or stored Pokemon",
			callback:    commandList,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Throw a Pokeball at a Pokemon",
			callback:    commandCatch,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the prev page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
