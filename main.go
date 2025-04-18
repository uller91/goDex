package main

import (
	"bufio"
	"os"
	"fmt"

	//"github.com/uller91/goDex/internal/apiInter"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/"
)

func main() {
	cfg := &config{}

	//url := baseUrl + "location-area/"
	//fmt.Println(apiInter.RequestLocation(url).Count)

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
	command, ok := getCommand()[input]

	if ok != true {
		fmt.Printf("Unknown command\n")
	} else {
		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error happened!\n")
		}
	}
	}
}
