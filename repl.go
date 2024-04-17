package main

import (
	"os"

	"github.com/ptenteromano/pokedexcli/internal/pokecache"
)

type config struct {
	cache   *pokecache.Cache
	mapUrl  string
	mapbUrl string
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
}
