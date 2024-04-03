package main

import (
	"os"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "-- Exits the program",
		callback: func() {
			os.Exit(0)
		},
	},
	"map": {
		name:        "map",
		description: "-- Shows the region areas of the pokemon world",
		callback: func() {
			pokeapi.GetLocations()
		},
	},
}
