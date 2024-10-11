package main

import (
	"time"

	"github.com/SumDeusVitae/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient         pokeapi.Client
	nextLocationsURL      *string
	prevLocationsURL      *string
	caughtPokemon         map[string]pokeapi.Pokemon
	history               []string
	current_history_index int
}

func main() {
	pokeClient := pokeapi.NewClient(time.Hour)
	cfg := &config{
		pokeapiClient:         pokeClient,
		caughtPokemon:         make(map[string]pokeapi.Pokemon),
		history:               []string{},
		current_history_index: -1,
	}
	startRepl(cfg)
}
