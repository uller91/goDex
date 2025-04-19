package main

import (
	"bufio"
	"os"
	"fmt"
	"time"

	"github.com/uller91/goDex/internal/cache"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/"
)

func main() {
	cfg := &config{Cache: cache.NewCache(5 * time.Minute),}
	//cache := cache.NewCache(5 * time.Minute) //interval of cleaning

	//cache.Add("now", []byte("123"))
	//val, _ := cache.Get("now")
	//fmt.Println(val)
	//val2, _ := cache.Get("not now")
	//fmt.Println(val2)
	//fmt.Println(cache.Data["now"])

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
		/*
		if command.name == "map" || command.name == "mapb" {
			err := command.callback(cfg, cache)
			if err != nil {
				fmt.Printf("Error happened!\n")
			}
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Printf("Error happened!\n")
			}
		}
			*/
	}
	}
}
