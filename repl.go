package main

import (
	"bufio"
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
			command.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type config struct {
	commands       map[string]cliCommand
	previousMapUrl string
	nextMapUrl     string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
