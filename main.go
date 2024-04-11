package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ptenteromano/pokedexcli/internal/pokecache"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex CLI!")

	conf := config{
		cache: pokecache.NewCache(20 * time.Second),
	}

	for {
		fmt.Print("pokedex > ")
		reader.Scan()
		command := reader.Text()

		if command == "help" {
			helpCommand.callback(&conf)
			continue
		}

		if res, ok := commands[command]; ok {
			res.callback(&conf)
		} else {
			fmt.Println("Command not found. Type 'help' to see all available commands.")
		}
	}
}
