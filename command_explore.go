package main

import (
	"fmt"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
)

var exploreCommand = cliCommand{
	name:        "explore",
	description: "-- Shows the pokemon in a location",
	callback:    callExploreLocation,
}

func callExploreLocation(_ *config, args ...string) {
	if len(args) == 0 {
		fmt.Println("No location provided. Type 'help' to see all available commands.")
	}

	// call the function from the pokeapi package
	pokeapi.ExploreLocation(args[0])
}
