package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(config *Config) error {
	var url string
	if config.Previous == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = config.Previous
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var locations Location
	decode := json.NewDecoder(res.Body)
	err = decode.Decode(&locations)
	if err != nil {
		return err
	}
	for _, data := range locations.Results {
		fmt.Println(data.Name)
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}
