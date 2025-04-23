package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.args) == 0 {
		return errors.New(
			"need a location name to explore. Usage: explore [location]",
		)
	}
	location := cfg.args[0]

	locationPokemonRes, err := cfg.pokeapiClient.ListLocationPokemon(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationPokemonRes.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}
