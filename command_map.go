package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Previous string         `json:"previous"`
	Next     string         `json:"next"`
	Results  []LocationArea `json:"results"`
}

func fetchLocationAreas(url string) ([]LocationArea, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var locationAreas LocationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&locationAreas)
	if err != nil {
		return nil, err
	}

	return locationAreas.Results, nil
}

func commandMap() error {
	locationAreas, err := fetchLocationAreas("")
	if err != nil {
		return err
	}

	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}

	return nil
}
