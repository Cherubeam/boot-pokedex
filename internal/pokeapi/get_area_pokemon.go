package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetAreaPokemon(locationArea string) (AreaPokemonResponse, error) {
	url := BaseURL + "location-area/" + locationArea

	if data, ok := c.cache.Get(url); ok {
		var result AreaPokemonResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			return AreaPokemonResponse{}, err
		}
		return result, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return AreaPokemonResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaPokemonResponse{}, err
	}

	if res.StatusCode > 299 {
		return AreaPokemonResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	c.cache.Add(url, body)

	var result AreaPokemonResponse

	if err = json.Unmarshal(body, &result); err != nil {
		return AreaPokemonResponse{}, err
	}

	return result, nil
}
