package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    callbackMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Return list of pokemons in the area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Trying to catch pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspects catched pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists pokemons in pokedex",
			callback:    callbackPokedex,
		},
		"history": {
			name:        "history",
			description: "Returns history of t commands",
			callback:    callbackHistory,
		},
		"clear": {
			name:        "clear",
			description: "Clears cli",
			callback:    callbackClear,
		},
	}

}
