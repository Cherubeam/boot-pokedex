# Boot Pokedex

A terminal-based Pokédex built in Go that lets you explore the Pokémon world directly from the command line.

## Overview

This project is a simple CLI app that interacts with the public Pokémon API to help you:

- browse location areas with `map` and `mapb`
- explore Pokémon in a specific area with `explore`
- catch Pokémon with `catch`
- inspect caught Pokémon with `inspect`
- view your collection using `pokedex`
- get help and command usage with `help`

## Features

- lightweight Go REPL interface
- region and area browsing through the PokéAPI
- Pokémon catch simulation
- personal Pokédex tracking
- inspection of stats, types, height, and weight

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/Cherubeam/boot-pokedex.git
   cd boot-pokedex
   ```

2. Run the app:
   ```bash
   go run .
   ```

3. Use the commands in the REPL:
   ```text
   Pokedex > help
   Pokedex > map
   Pokedex > explore kanto
   Pokedex > catch pikachu
   Pokedex > inspect pikachu
   Pokedex > pokedex
   Pokedex > exit
   ```

## Available Commands

- `help` — display available commands
- `map` — show the next page of location areas
- `mapb` — show the previous page of location areas
- `explore <area>` — list Pokémon found in an area
- `catch <pokemon>` — attempt to catch a Pokémon
- `inspect <pokemon>` — show details for a caught Pokémon
- `pokedex` — show all caught Pokémon
- `exit` — quit the app

## Tech Stack

- Go
- Pokémon API

## License

This project is for educational and personal use.
