package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		textArr := cleanInput(text)
		if len(textArr) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", textArr[0])
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return []string{}
	}

	words := strings.Fields(text)

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}
