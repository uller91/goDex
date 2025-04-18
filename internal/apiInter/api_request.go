package apiInter

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"

	"github.com/uller91/goDex/internal/cache"
)

func normalizeURLLocationArea(url string) string {
    if url == "https://pokeapi.co/api/v2/location-area/" {
        return "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
    }
    return url
}

func RequestData(url string, cache *cache.Cache) []byte {
	//check the cache
	if data, found := cache.Get(url); found {
		fmt.Println("Extracting data from cache...")
		return data
	}

	//request for a data
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	cache.Add(url, data)
	//fmt.Printf("%s", body)
	return data
}

func RequestLocation(url string, cache *cache.Cache) LocationResults {
	url = normalizeURLLocationArea(url)
	data := RequestData(url, cache)

	//unmarshal data
	locationsList := LocationResults{}
	err := json.Unmarshal(data, &locationsList)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return locationsList
}

func RequestPokemon(url string, cache *cache.Cache) PokemonResults {
	data := RequestData(url, cache)

	//unmarshal data
	pokemonList := PokemonResults{}
	err := json.Unmarshal(data, &pokemonList)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return pokemonList
}

func RequestPokemonStats(url string, cache *cache.Cache) PokemonStats {
	data := RequestData(url, cache)

	//unmarshal data
	pokemonStats := PokemonStats{}
	err := json.Unmarshal(data, &pokemonStats)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return pokemonStats
}