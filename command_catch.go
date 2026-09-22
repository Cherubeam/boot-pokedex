package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch command requires a Pokemon name")
	}

	name := args[0]
	res, err := cfg.apiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	catchChance := rand.Intn(500)
	if catchChance < res.BaseExperience {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[name] = res
	}
	return nil
}
