package main

import (
	"fmt"

	"github.com/Cherubeam/boot-pokedex/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
	if cfg.mapState.nextURL == nil {
		initialURL := pokeapi.BaseURL + "location-area/"
		cfg.mapState.nextURL = &initialURL
	}

	res, err := cfg.apiClient.GetLocationAreas(*cfg.mapState.nextURL)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.mapState.nextURL = res.Next
	cfg.mapState.previousURL = res.Previous

	return nil
}
