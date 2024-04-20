package commands

import (
	"os"

	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
	"github.com/ptenteromano/pokedexcli/internal/pokecache"
)

type Config struct {
	Cache   *pokecache.Cache // All cached data stored here
	mapUrl  string           // These have nothing to do with the Cache - this just stores "where" the user is in the map
	mapbUrl string
	pokemon []*pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	Callback    func(c *Config, args ...string)
}

var AllCmds = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "-- Exits the program",
		Callback: func(_ *Config, _ ...string) {
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

func (c Config) caughtPokemon(pokemon string) bool {
	for _, p := range c.pokemon {
		if p.Name == pokemon {
			return p.Caught
		}
	}

	return false
}

func (c Config) storedPokemon(pokemon string) *pokeapi.Pokemon {
	for _, p := range c.pokemon {
		if p.Name == pokemon {
			return p
		}
	}

	return nil
}

func (c Config) listCaughtPokemon() []string {
	var caught []string
	for _, p := range c.pokemon {
		if p.Caught {
			caught = append(caught, p.Name)
		}
	}

	return caught
}
