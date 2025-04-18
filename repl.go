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

		fmt.Printf("Your command was: %s\n", textArr[0])
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	textArr := strings.Fields(lower)
	return textArr
}
