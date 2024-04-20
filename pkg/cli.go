package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ptenteromano/pokedexcli/internal/commands"
	"github.com/ptenteromano/pokedexcli/internal/pokecache"
)

func Run() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex CLI!")

	conf := commands.Config{
		Cache: pokecache.NewCache(20 * time.Second),
	}

	for {
		fmt.Print("pokedex > ")
		reader.Scan()
		command := reader.Text()

		if command == "help" {
			commands.HelpCommand.Callback(&conf)
			continue
		}

		// Split arguments by whitespace
		args := strings.Split(command, " ")

		if res, ok := commands.AllCmds[args[0]]; ok {
			res.Callback(&conf, args[1:]...)
		} else {
			fmt.Println("Command not found. Type 'help' to see all available commands.")
		}
	}
}
