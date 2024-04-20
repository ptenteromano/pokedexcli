package commands

import "fmt"

var pokedex = cliCommand{
	name:        "pokedex",
	description: "-- Shows the pokedex",
	Callback:    callPokedex,
}

func callPokedex(c *Config, args ...string) {
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
