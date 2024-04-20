package main

import "fmt"

var pokedex = cliCommand{
	name:        "pokedex",
	description: "-- Shows the pokedex",
	callback:    callPokedex,
}

func callPokedex(c *config, args ...string) {
	if len(c.pokemon) == 0 {
		fmt.Println("No pokemon caught yet - go get some!")
		return
	}

	caughtPokemon := c.listCaughtPokemon()

	fmt.Println("Your Pokedex:")
	for _, p := range caughtPokemon {
		fmt.Println("  -", p)
	}
}
