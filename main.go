package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner error: %v", err)
		}

		input := scanner.Text()
		clean := cleanInput(input)

		command, ok := commands[clean[0]]
		if ok {
			command.callback()
		} else {
			fmt.Println("Unknown command")
		}

	}
}
