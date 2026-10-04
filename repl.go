package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	s = strings.TrimSpace(s)

	candidates := strings.Split(s, " ")
	var result []string

	for _, c := range candidates {
		if c != "" && c != " " {
			result = append(result, c)
		}
	}

	return result
}

type CommandCallback func() error

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
