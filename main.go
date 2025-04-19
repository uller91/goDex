package main

import (
	"bufio"
	"os"
	"fmt"
	"time"

	"github.com/uller91/goDex/internal/cache"
	"github.com/uller91/goDex/internal/apiInter"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/"
)

func main() {
	pokedex := make(map[string]apiInter.PokemonStats)
	cfg := &config{Cache: cache.NewCache(5 * time.Minute), Pokedex: pokedex}
	prm := &parameters{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
    	err := scanner.Err()
    	if err != nil {
        	fmt.Errorf("Error %v!", err)
    }
	usrInput := scanner.Text()
	usrInputClean := cleanInput(usrInput)

	input := usrInputClean[0]
	if input == "explore" && len(usrInputClean)<2 {
		fmt.Printf("Enter the area to explore!\n")
		continue
	}
	if input == "catch" && len(usrInputClean)<2 {
		fmt.Printf("Enter the pokemon to catch!\n")
		continue
	}

	if input == "inspect" && len(usrInputClean)<2 {
		fmt.Printf("Enter the pokemon to inspect!\n")
		continue
	}

	if input == "explore" || input == "catch" || input == "inspect" {
		prm.Key = usrInputClean[1]
	}
	
	command, ok := getCommand()[input]

	if ok != true {
		fmt.Printf("Unknown command\n")
		continue
	} else {
		err := command.callback(cfg, prm)
		if err != nil {
			fmt.Printf("Error happened!\n")
		}
	}
	}
}
