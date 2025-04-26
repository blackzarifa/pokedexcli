package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience
	if baseExp <= 0 {
		baseExp = 1
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := 50.0 / float64(baseExp)
	catch := rng.Float64() < chance

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if catch {
		fmt.Println(pokemonName + " was caught!")
		cfg.pokedex[pokemonName] = pokemon
		return nil
	}

	fmt.Println(pokemonName + " escaped!")
	return nil
}
