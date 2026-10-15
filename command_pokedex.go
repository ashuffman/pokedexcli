package main

import (
	"fmt"
)

func commandPokedex(cfg *config, par string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex.Dex {
		fmt.Println(" - ", pokemon.Name)
	}

	return nil
}
