package main

import (
	"strings"

	"github.com/uller91/goDex/internal/cache"
	"github.com/uller91/goDex/internal/apiInter"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *parameters) error
}

type config struct {
	Cache		*cache.Cache
	Next		*string
	Previous	*string
	Pokedex		map[string]apiInter.PokemonStats
}

type parameters struct {
	Key	string
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:		 "map",
			description: "Display next 20 locations in the world",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Display previous 20 locations in the world",
			callback:	 commandMapb,
		},
		"explore": {
			name:		 "explore <area_name>",
			description: "Display all the pokemon found in <area_name>",
			callback:	 commandExplore,
		},
		"catch": {
			name:		 "catch <pokemon>",
			description: "Try to catch the <pokemon>",
			callback:	 commandCatch,
		},
		"inspect": {
			name:		 "inspect <pokemon>",
			description: "Inspect the <pokemon>",
			callback:	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "Display all the pokemons caught!",
			callback:	 commandPokedex,
		},
}
}