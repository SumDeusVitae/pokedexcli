<div align="center">
  <h1>Pokedex REPL<br>(Read, Evaluate, Print, and Loop)</h1>
</div>


## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Credit](#credit)
- [License](#license)

## Introduction
A REPL, or Read-Eval-Print Loop, is a simple interactive programming environment that takes user input, evaluates it, and returns the result to the user. This project is about a Pokedex-like REPL in Go that uses the PokeAPI to fetch data about  Pokemon.

## Installation
1. Clone the repository:
```bash
 git clone https://github.com/SumDeusVitae/pokedexcli.git
```

2. Navigate to the Folder:
```bash
  cd pokedexcli
 ```

3. Build the Application:
```bash
  go build
 ```


## Usage
To run the project, use the following command:
```bash
./pokedexcli
```


## Commands 
- `help`:                     Displays a help message
- `exit`:                     Exit the Pokedex
- `clear`:                    Clears cli
- `map`:                      Get the next page of locations
- `mapb`:                     Get the previous page of locations
- `inspect {pokemon_name}`:   Inspects catched pokemon
- `explore {location_area}`:  Return list of pokemons in the area
- `catch {pokemon_name}`:     Trying to catch pokemon
- `pokedex`:                  Lists pokemons in pokedex
- `history`:                  Returns history of commands

## Credit
> [**chzyer**](https://github.com/chzyer/readline) >> for pure go(golang) implementation of GNU-Readline kind library 



## License
This project is licensed under the [MIT License](LICENSE).
