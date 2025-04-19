package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		textArr := cleanInput(scanner.Text())
		if len(textArr) == 0 {
			continue
		}

		cmdName := textArr[0]
		cmd, ok := getCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	textArr := strings.Fields(lower)
	return textArr
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous 20 location areas",
			callback:    commandMapB,
		},
		"map": {
			name:        "map",
			description: "Shows next 20 location areas",
			callback:    commandMap,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
