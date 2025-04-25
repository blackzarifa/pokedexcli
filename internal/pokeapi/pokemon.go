package pokeapi

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	return fetchFromAPI[Pokemon](c, url)
}
