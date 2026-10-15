package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ashuffman/pokedexcli/internal/pokeapi"
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

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner error: %v", err)
		}

		input := scanner.Text()
		clean := cleanInput(input)

		command, ok := cfg.commands[clean[0]]
		if ok {
			var par string
			if len(clean) > 1 {
				par = clean[1]
			} else {
				par = ""
			}
			command.callback(cfg, par)
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type config struct {
	commands       map[string]cliCommand
	pokedex        pokeapi.Pokedex
	previousMapUrl *string
	nextMapUrl     *string
	client         pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
