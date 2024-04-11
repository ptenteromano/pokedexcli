package main

import (
	"os"
)

type config struct {
	mapUrl  string
	mapbUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config)
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "-- Exits the program",
		callback: func(_ *config) {
			os.Exit(0)
		},
	},
	"map":  mapCommand,
	"mapb": mapbCommand,
}
