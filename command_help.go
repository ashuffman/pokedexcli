package main

import "fmt"

func commandHelp(cfg *config, par string) error {
	fmt.Print(`
-----------------------
Welcome to the Pokedex!
Usage:

`)

	commands := cfg.commands
	for c := range commands {
		fmt.Printf("%v: %v \n", commands[c].name, commands[c].description)
	}

	fmt.Print("\n \n")

	return nil
}
