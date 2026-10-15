package main

import (
	"time"

	"github.com/ashuffman/pokedexcli/internal/pokeapi"
)

func main() {
	const (
		cacheInterval = 5 * time.Second
	)

	cfg := config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Displays the names of 20 location areas in the Pokemon world.",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Displays the previous 20 location areas in the Pokemon world.",
				callback:    commandMapB,
			},
			"explore": {
				name:        "explore",
				description: "Displays a list of potential Pokémon encounters for a specified location.",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Attempts to catch specified Pokémon. Caught Pokémon are added to user Pokédex.",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Displays information about a specified Pokémon from your Pokédex",
				callback:    commandInspect,
			},
		},
		pokedex: pokeapi.NewPokedex(),
		client:  pokeapi.NewClient(cacheInterval),
	}

	startRepl(&cfg)
}
