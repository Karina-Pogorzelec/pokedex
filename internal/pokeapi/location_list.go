package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {

	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	if data, ok := c.cache.Get(fullURL); ok {
		var locationsResp RespShallowLocations
    	if err := json.Unmarshal(data, &locationsResp); err != nil {
     		return RespShallowLocations{}, err
    	}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(fullURL, body)

	var locationsResp RespShallowLocations
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil

}