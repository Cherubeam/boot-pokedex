package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Cherubeam/boot-pokedex/internal/pokeapi"
)

type mapState struct {
	nextURL     *string
	previousURL *string
}

type config struct {
	commands  map[string]cliCommand
	apiClient *pokeapi.Client
	mapState  mapState
	pokedex   map[string]pokeapi.PokemonResponse
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !reader.Scan() {
			break
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

	if err := reader.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Catches a Pokemon and adds it to the user's Pokedex",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokemon located in a specified area.",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the name, height, weight, stats, and type(s) of a Pokemon.",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the names of the Pokemon the user has caught.",
			callback:    commandPokedex,
		},
	}
}
