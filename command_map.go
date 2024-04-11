package main

import (
	"github.com/ptenteromano/pokedexcli/internal/pokeapi"
)

var mapCommand = cliCommand{
	name:        "map",
	description: "-- Shows the region areas of the pokemon world",
	callback: func(c *config) {
		callMapLocations(c, false)
	},
}

var mapbCommand = cliCommand{
	name:        "mapb",
	description: "-- Shows the previous region areas of the pokemon world",
	callback: func(c *config) {
		callMapLocations(c, true)
	},
}

func callMapLocations(c *config, back bool) {
	url := c.mapUrl

	if back {
		url = c.mapbUrl
	}

	areas, next, previous := pokeapi.GetLocations(url)

	c.mapUrl = next
	c.mapbUrl = previous

	for _, area := range areas {
		println(area.Name)
	}
}
