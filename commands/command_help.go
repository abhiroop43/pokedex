package commands

import "fmt"

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("The available commands are as follows:")
	fmt.Println("help: Displays this guide")
	fmt.Println("exit: Closes the Pokedex")
	fmt.Println("")
	return nil
}
