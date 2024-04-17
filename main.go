package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		// Split arguments by whitespace
		args := strings.Split(command, " ")

		if res, ok := commands[args[0]]; ok {
			res.callback(&conf, args[1:]...)
		} else {
			fmt.Println("Command not found. Type 'help' to see all available commands.")
		}
	}
}
