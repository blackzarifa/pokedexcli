package pokeapi

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	return fetchFromAPI[LocationAreaResponse](c, url)
}

func (c *Client) ListLocationPokemon(
	areaName string,
) (LocationAreaDetailsResponse, error) {
	url := baseURL + "/location-area/" + areaName
	return fetchFromAPI[LocationAreaDetailsResponse](c, url)
}
