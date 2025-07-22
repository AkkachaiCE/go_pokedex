package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(config *Config) error {
	var url string
	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = config.Next
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

type Location struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []LocationRef `json:"results"`
}

type LocationRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
