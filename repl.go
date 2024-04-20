package main

import (
	"os"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
	"github.com/ptenteromano/pokedexcli/internal/pokecache"
)

type config struct {
	cache   *pokecache.Cache // All cached data stored here
	mapUrl  string           // These have nothing to do with the Cache - this just stores "where" the user is in the map
	mapbUrl string
	pokemon []*pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string)
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "-- Exits the program",
		callback: func(_ *config, _ ...string) {
			os.Exit(0)
		},
	},
	"map":     mapCommand,
	"mapb":    mapbCommand,
	"explore": exploreCommand,
	"catch":   catchCommand,
	"inspect": inspectCommand,
	"pokedex": pokedex,
}

func (c config) caughtPokemon(pokemon string) bool {
	for _, p := range c.pokemon {
		if p.Name == pokemon {
			return p.Caught
		}
	}

	return false
}

func (c config) storedPokemon(pokemon string) *pokeapi.Pokemon {
	for _, p := range c.pokemon {
		if p.Name == pokemon {
			return p
		}
	}

	return nil
}

func (c config) listCaughtPokemon() []string {
	var caught []string
	for _, p := range c.pokemon {
		if p.Caught {
			caught = append(caught, p.Name)
		}
	}

	return caught
}
