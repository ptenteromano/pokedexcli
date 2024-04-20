package main

import (
	"fmt"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
)

var catchCommand = cliCommand{
	name:        "catch",
	description: "-- Attempts to catch a pokemon",
	callback:    callCatchPokemon,
}

func callCatchPokemon(_ *config, args ...string) {
	if len(args) == 0 {
		fmt.Println("No pokemon provided. Type 'help' to see all available commands.")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	// call the function from the pokeapi package
	success, err := pokeapi.AttemptCatch(args[0])

	if err != nil {
		fmt.Println("Error:", err)
	}

	if success {
		fmt.Printf("%s was caught!\n", args[0])
	} else {
		fmt.Printf("Oh no! %s broke free!\n", args[0])
	}
}
