package main

import (
	"errors"
	"fmt"
	"time"

	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	_, ok := cfg.caughtPokemon[pokemonName]
	if ok {
		return fmt.Errorf("You already caught: %s", pokemonName)
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	// Create a Rand instance
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random integer between 0 and n
	n := pokemonResp.BaseExperience
	fmt.Printf("Base experiense of %s is: %d\n", pokemonResp.Name, n)
	fmt.Println("You are trying to catch pokemon...")
	if r.Intn(n) > 20 {
		return fmt.Errorf("You did not catch - %s", pokemonResp.Name)
	}
	cfg.caughtPokemon[pokemonName] = pokemonResp
	fmt.Printf("You caught - %s, you can inspect it now.\n", pokemonResp.Name)

	return nil
}
