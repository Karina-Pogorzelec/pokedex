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

	var locationsResp RespShallowLocations
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil

}