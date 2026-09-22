package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
	if data, ok := c.cache.Get(url); ok {
		var result LocationAreaResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return result, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	c.cache.Add(url, body)

	var result LocationAreaResponse

	if err = json.Unmarshal(body, &result); err != nil {
		return LocationAreaResponse{}, err
	}

	return result, nil
}
