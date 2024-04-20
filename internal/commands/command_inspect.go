package commands

import "fmt"

var inspectCommand = cliCommand{
	name:        "inspect",
	description: "-- Shows the details of a pokemon",
	Callback:    callInspectPokemon,
}

func callInspectPokemon(c *Config, args ...string) {
	if len(args) == 0 {
		fmt.Println("No pokemon provided. Type 'help' to see all available commands.")
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments provided. Type 'help' to see all available commands.")
	}

	if !c.caughtPokemon(args[0]) {
		fmt.Printf("You don't have %s, and can't inspect it!\n", args[0])
		return
	}

	for _, p := range c.pokemon {
		if p.Name == args[0] {
			fmt.Printf("Name: %s\n", p.Name)
			fmt.Printf("Height: %d\n", p.Height)
			fmt.Printf("Weight: %d\n", p.Weight)
			fmt.Println("Stats:")
			for _, s := range p.Stats {
				fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
			}
			fmt.Println("Types:")
			for _, t := range p.Types {
				fmt.Printf("  - %s\n", t.Type.Name)
			}
			return
		}
	}
}
