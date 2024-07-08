package commands

import "fmt"

func CommandHelp() error {
	fmt.Printf("Welcome to the Pokedex\n")
	fmt.Printf("The available commands are as follows:\n\n")

	availableCommands := GetCommands()

	for _, command := range availableCommands {
		fmt.Printf("%s: %s\n\n", command.Name, command.Description)
	}
	return nil
}
