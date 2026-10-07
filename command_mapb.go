package main

import (
	"fmt"

	"github.com/ashuffman/pokedexcli/internal/pokeapi"
)

func commandMapB(cfg *config) error {
	// if we're on the first page, return the "first page" message and return nil
	if cfg.previousMapUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	// set url to previous
	url := *cfg.previousMapUrl

	mapResponse, err := pokeapi.MakeRequest(url)
	if err != nil {
		return err
	}

	cfg.nextMapUrl = mapResponse.Next
	cfg.previousMapUrl = mapResponse.Previous

	for _, location := range mapResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
