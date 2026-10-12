package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	// set url to location-area end-point
	url := "https://pokeapi.co/api/v2/location-area/"

	// if cfg.nextMapURL is not nil dereference it and assign it to the url variable
	if cfg.nextMapUrl != nil {
		url = *cfg.nextMapUrl
	}

	mapResponse, err := cfg.client.MakeRequest(url)
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
