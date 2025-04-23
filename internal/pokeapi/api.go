package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetchFromAPI[T any](c *Client, url string) (T, error) {
	var result T

	if cachedData, found := c.cache.Get(url); found {
		err := json.Unmarshal(cachedData, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
