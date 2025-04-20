package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, found := c.cache.Get(url); found {
		locationsResp := LocationAreaResponse{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, data)

	locationsResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationsResp, nil
}
