package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcomt to Pokedex! \nUsage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
