package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) GetLocation(locationName string) (PokemonLocationData, error) {
	fullURL := baseURL + "/location-area/" + locationName

	if data, ok := c.cache.Get(fullURL); ok {
		var location PokemonLocationData
    	if err := json.Unmarshal(data, &location); err != nil {
     		return PokemonLocationData{}, err
    	}
		return location, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonLocationData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonLocationData{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonLocationData{}, err
	}

	c.cache.Add(fullURL, body)

	var location PokemonLocationData
	err = json.Unmarshal(body, &location)
	if err != nil {
		return PokemonLocationData{}, err
	}

	return location, nil	
}