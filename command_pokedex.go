package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) < 1 {
		return errors.New("you have not caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range cfg.pokedex {
		fmt.Println("  - " + p.Name)
	}

	return nil
}
