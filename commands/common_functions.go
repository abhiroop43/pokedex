package commands

import "abhiroop43/pokedex/dto"

func GetCommands() map[string]dto.CliCommand {
	return map[string]dto.CliCommand{
		"help": {
			Name:        "help",
			Description: "Prints the help menu",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Closes the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Shows the next 20 locations on the map",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous 20 locations on the map",
			Callback:    CommandMapB,
		},
	}
}
