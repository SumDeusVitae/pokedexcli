package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught %s", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, each := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", each.Stat.Name, each.BaseStat)
	}
	fmt.Println("Types:")
	for _, s := range pokemon.Types {
		fmt.Printf("	- %s\n", s.Type.Name)
	}

	return nil
}
