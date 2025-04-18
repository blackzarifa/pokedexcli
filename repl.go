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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
