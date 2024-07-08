package main

import (
	"abhiroop43/pokedex/commands"
	"abhiroop43/pokedex/dto"
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

func startRepl(config *dto.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := commands.GetCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.Callback(config)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func cleanInput(str string) []string {
	lowerString := strings.ToLower(str)
	words := strings.Fields(lowerString)
	return words
}
