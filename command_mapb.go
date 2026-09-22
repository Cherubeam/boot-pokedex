package main

import (
	"fmt"
)

func commandMapb(cfg *config, args []string) error {
	if cfg.mapState.previousURL == nil {
		return fmt.Errorf("no previous page")
	}

	res, err := cfg.apiClient.GetLocationAreas(*cfg.mapState.previousURL)
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
