package main

import (
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
