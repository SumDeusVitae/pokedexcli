package main

import (
	"time"

	"github.com/SumDeusVitae/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Hour)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
