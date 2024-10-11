package main

import "fmt"

func callbackPokedex(cfg *config, args ...string) error {
	catchedPokemons := cfg.caughtPokemon

	if len(catchedPokemons) < 1 {
		return fmt.Errorf("Your pokedex empty")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range catchedPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
