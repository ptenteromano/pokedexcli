package commands

import (
	"fmt"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
)

var exploreCommand = cliCommand{
	name:        "explore",
	description: "-- Shows the pokemon in a location",
	Callback:    callExploreLocation,
}

func callExploreLocation(_ *Config, args ...string) {
	if len(args) == 0 {
		fmt.Println("No location provided. Type 'help' to see all available commands.")
	}

	// call the function from the pokeapi package
	pokemon, err := pokeapi.ExploreLocation(args[0])

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	for _, p := range pokemon {
		fmt.Println("- " + p)
	}
}
