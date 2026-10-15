package main

import (
	"fmt"
)

func commandInspect(cfg *config, par string) error {
	pokemon, ok := cfg.pokedex.Dex[par]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf(
		`
Name: %v
Height: %v
Weight: %v
Stats:	
`,
		pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stat {
		fmt.Printf("    -%v: %v \n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, t := range pokemon.Types {
		fmt.Printf("    - %v \n", t.Type.Name)
	}

	return nil
}
