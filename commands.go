package main

import (
	"os"
	"fmt"
	
	"github.com/uller91/goDex/internal/apiInter"
)

func commandExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommand() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	url := baseUrl + "location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	locationList := apiInter.RequestLocation(url, cfg.Cache)
	cfg.Previous = &locationList.Previous
	cfg.Next = &locationList.Next
	//fmt.Println(cfg.Next)
	//fmt.Println(*cfg.Next)
	//fmt.Println(*cfg.Previous)
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	url := baseUrl + "location-area/"
	if cfg.Previous != nil && *cfg.Previous != "" {
		url = *cfg.Previous
	}
	locationList := apiInter.RequestLocation(url, cfg.Cache)
	cfg.Previous = &locationList.Previous
	cfg.Next = &locationList.Next
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}
	return nil
}