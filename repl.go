package main

import (
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next		*string
	Previous	*string
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
}

}