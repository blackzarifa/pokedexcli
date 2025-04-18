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

var commands = map[string]cliCommand{}

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
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

func commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
