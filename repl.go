package main

import (
	"abhiroop43/pokedex/commands"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]

		switch command {
		case "help":
			fmt.Println("Welcome to the Pokedex")
			fmt.Println("The available commands are as follows:")
			fmt.Println("help: Displays this guide")
			fmt.Println("exit: Closes the Pokedex")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func cleanInput(str string) []string {
	lowerString := strings.ToLower(str)
	words := strings.Fields(lowerString)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    commands.CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes the Pokedex",
			callback:    commands.CommandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 locations on the map",
			callback:    nil,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 locations on the map",
			callback:    nil,
		},
	}
}
