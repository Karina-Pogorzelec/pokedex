package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(fullURL); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("Pokemon not found")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, body)

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil

}