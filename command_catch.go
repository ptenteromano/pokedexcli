package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
)

var catchCommand = cliCommand{
	name:        "catch",
	description: "-- Attempts to catch a pokemon",
	callback:    callCatchPokemon,
}

func callCatchPokemon(c *config, args ...string) {
	if len(args) == 0 {
		fmt.Println("No pokemon provided. Type 'help' to see all available commands.")
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments provided. Type 'help' to see all available commands.")
		return
	}

	name := args[0]

	if c.caughtPokemon(name) {
		fmt.Printf("You already have %s!\n", name)
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	var pokemon *pokeapi.Pokemon
	var err error

	// Check if the pokemon is already cached to avoid redundant API calls
	if c.storedPokemon(name) != nil {
		pokemon = c.storedPokemon(name)
	} else {
		pokemon, err = pokeapi.FetchPokemonData(name)

		if err != nil {
			fmt.Println("Error:", err)
		}

		c.pokemon = append(c.pokemon, pokemon)
	}

	chance := rand.IntN(350) // Mewtwo is 340

	if chance < pokemon.BaseExperience {
		fmt.Printf("Oh no! %s broke free!\n", name)
		return
	}

	pokemon.Caught = true
	fmt.Printf("%s was caught!\n", name)
	fmt.Println("You may now inspect it with the inspect command.")
}
