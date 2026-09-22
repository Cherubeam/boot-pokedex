package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("pokedex command does not take any arguments")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
