package main

import (
	"bufio"
	"fmt"
	"os"
)

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
