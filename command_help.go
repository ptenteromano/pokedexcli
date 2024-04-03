package main

import "fmt"

var helpCommand = cliCommand{
	name:        "help",
	description: "-- Lists all the commands available",
	callback:    help,
}

func help() {
	for _, cmd := range commands {
		if cmd.name == "help" {
			continue
		}

		fmt.Println(cmd.name, "\t", commands[cmd.name].description)
	}
}
