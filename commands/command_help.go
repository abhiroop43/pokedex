package commands

import (
	"abhiroop43/pokedex/dto"
	"fmt"
)

func CommandHelp(config *dto.Config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex\n")
	fmt.Printf("The available commands are as follows:\n\n")

	availableCommands := GetCommands()

	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", command.Name, command.Description)
	}
	fmt.Println("")
	return nil
}
