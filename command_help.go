package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}
