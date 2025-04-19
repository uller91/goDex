package main

import (
	"os"
	"fmt"
	"math/rand"

	"github.com/uller91/goDex/internal/apiInter"
)

func commandExit(cfg *config, prm *parameters) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, prm *parameters) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommand() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, prm *parameters) error {
	url := baseUrl + "location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	locationList := apiInter.RequestLocation(url, cfg.Cache)
	cfg.Previous = &locationList.Previous
	cfg.Next = &locationList.Next
	//fmt.Println(cfg.Next)
	//fmt.Println(*cfg.Next)
	for _, location := range locationList.Results {
		fmt.Printf("- %v\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config, prm *parameters) error {
	url := baseUrl + "location-area/"
	if cfg.Previous != nil && *cfg.Previous != "" {
		url = *cfg.Previous
	}
	locationList := apiInter.RequestLocation(url, cfg.Cache)
	cfg.Previous = &locationList.Previous
	cfg.Next = &locationList.Next
	for _, location := range locationList.Results {
		fmt.Printf("- %v\n", location.Name)
	}
	return nil
}

func commandExplore(cfg *config, prm *parameters) error {
	fmt.Printf("Exploring %v...\n", prm.Key)
	url := baseUrl + "location-area/" + prm.Key

	pokemonList := apiInter.RequestPokemon(url, cfg.Cache)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList.Results {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, prm *parameters) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", prm.Key)
	url := baseUrl + "pokemon/" + prm.Key

	pokemonStats := apiInter.RequestPokemonStats(url, cfg.Cache)
	//fmt.Printf("Data about %v was found!\n", pokemonStats.Name)

	//random chance to catch
	randomVal := rand.Intn(500)
	if randomVal > pokemonStats.BaseExperience && prm.Key == pokemonStats.Name {
		fmt.Printf("%v was caught!\n", pokemonStats.Name)
		cfg.Pokedex[pokemonStats.Name] = pokemonStats
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%v escaped!\n", pokemonStats.Name)
	}

	return nil
}

func commandInspect(cfg *config, prm *parameters) error {
	pokemon, caught := cfg.Pokedex[prm.Key]
	if !caught {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Printf(" - %v\n", tp.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, prm *parameters) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}