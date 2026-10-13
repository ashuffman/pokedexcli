package main

import (
	"fmt"

	"github.com/ashuffman/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, par string) error {
	// set url to area-location endpoint + par (the passed parameter)
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	url := baseURL + par

	// make http request and store resulting []byte in data
	data, err := cfg.client.MakeRequest(url)
	if err != nil {
		return err
	}

	// marshal the response into a ExplorationEncounters struct
	exploreResponse, err := pokeapi.ExploreUnmarshal(data)
	if err != nil {
		return err
	}

	// print the name of each pokemon in the list of Pokemon encounters
	for _, p := range exploreResponse.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}

	return nil
}
