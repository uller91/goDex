package apiInter

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func RequestLocation(url string) LocationResults {
	//request for a data
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	//fmt.Printf("%s", body)

	//unmarshal data
	locationsList := LocationResults{}
	err = json.Unmarshal(body, &locationsList)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return locationsList
}
