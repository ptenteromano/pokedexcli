package main

import (
	"encoding/json"
	"fmt"

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
	var url string

	if back {
		url = c.mapbUrl
	} else {
		url = c.mapUrl
	}

	entry, ok := c.cache.Get(url)

	var areas pokeapi.LocationResponse

	if ok {
		json.Unmarshal(entry, &areas)
	} else {
		areas = pokeapi.GetLocations(url)
		byteData, _ := json.Marshal(areas)

		c.cache.Add(url, byteData)
	}

	c.mapUrl = areas.Next
	c.mapbUrl = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
}
