package commands

import "fmt"

var HelpCommand = cliCommand{
	name:        "help",
	description: "-- Lists all the commands available",
	Callback:    help,
}

func help(_ *Config, _ ...string) {
	for _, cmd := range AllCmds {
		if cmd.name == "help" {
			continue
		}

		fmt.Println(cmd.name, "\t", AllCmds[cmd.name].description)
	}
}
