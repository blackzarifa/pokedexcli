package main

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Results []LocationArea `json:"results"`
}

var apiUrl string = "https://pokeapi.co/api/v2/"

func fetchLocationAreas() ([]LocationArea, error) {
	res, err := http.Get(apiUrl + "location-area")
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
