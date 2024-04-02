package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

var helpCommand = cliCommand{
	name:        "help",
	description: "-- Lists all the commands available",
	callback:    help,
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "-- Exits the program",
		callback: func() {
			os.Exit(0)
		},
	},
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex CLI!")

	for {
		fmt.Print("pokedex > ")
		reader.Scan()
		command := reader.Text()

		if command == "help" {
			helpCommand.callback()
			continue
		}

		if res, ok := commands[command]; ok {
			res.callback()
		} else {
			fmt.Println("Command not found. Type 'help' to see all available commands.")
		}
	}
}

func help() {
	for _, cmd := range commands {
		if cmd.name == "help" {
			continue
		}

		fmt.Println(cmd.name, "\t", commands[cmd.name].description)
	}
}
