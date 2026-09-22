package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	url := BaseURL + "pokemon/" + name + "/"

	if data, ok := c.cache.Get(url); ok {
		var result PokemonResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			return PokemonResponse{}, err
		}
		return result, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	if res.StatusCode > 299 {
		return PokemonResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	c.cache.Add(url, body)

	var result PokemonResponse

	if err = json.Unmarshal(body, &result); err != nil {
		return PokemonResponse{}, err
	}

	return result, nil
}
