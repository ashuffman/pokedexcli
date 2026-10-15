package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/ashuffman/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, par string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", par)

	// set url to pokemon endpoint + pokemon name
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	url := baseURL + par

	// make http request and store resulting []byte in data
	data, err := cfg.client.MakeRequest(url)
	if err != nil {
		return err
	}

	// marshal the response into a Pokemon struct
	pokemon, err := pokeapi.CatchUnmarshal(data)
	if err != nil {
		return err
	}

	// define catching logic?
	catchEffectiveness := rand.IntN(1000)
	if catchEffectiveness >= pokemon.BaseExperience {
		cfg.pokedex.Dex[pokemon.Name] = pokemon
		fmt.Printf("%v was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
