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

var (
	previousURL string
	nextURL     string
)

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

	previousURL = locationAreas.Previous
	nextURL = locationAreas.Next

	return locationAreas.Results, nil
}

func commandMap(cfg *config) error {
	locationAreas, err := fetchLocationAreas(nextURL)
	if err != nil {
		return err
	}

	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if previousURL == "" {
		return fmt.Errorf("you're on the first page")
	}

	locationAreas, err := fetchLocationAreas(previousURL)
	if err != nil {
		return err
	}

	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}

	return nil
}
