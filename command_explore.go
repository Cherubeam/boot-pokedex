package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore command requires a location area")
	}

	area := args[0]
	res, err := cfg.apiClient.GetAreaPokemon(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
