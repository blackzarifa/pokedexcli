package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		textArr := cleanInput(scanner.Text())
		if len(textArr) == 0 {
			continue
		}

		c := textArr[0]
		_, ok := commands[c]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		commands[c].callback()
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	textArr := strings.Fields(lower)
	return textArr
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
