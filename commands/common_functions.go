package commands

import "abhiroop43/pokedex/dto"

func GetCommands() map[string]dto.CliCommand {
	return map[string]dto.CliCommand{
		"help": {
			Name:        "help",
			Description: "Prints the help menu",
			Callback:    CommandHelp,
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
		"explore": {
			Name:        "explore (a location)",
			Description: "Lists all the Pokemons in a location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch (a pokemon)",
			Description: "Try to catch a Pokemon for your Pokedex",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect (a pokemon)",
			Description: "Inspect a Pokemon that has been caught and is available in your Pokedex",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "View your Pokedex",
			Description: "View your Pokedex and all the Pokemons caught",
			Callback:    CommandPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Closes the Pokedex",
			Callback:    CommandExit,
		},
	}
}
