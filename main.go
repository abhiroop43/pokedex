package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')
		switch text {
		case "help\n":
			err := commandHelp()
			if err != nil {
				fmt.Println("Error:", err.Error())
				os.Exit(1)
			}
		case "exit\n":
			err := commandExit()
			if err != nil {
				fmt.Println("Error:", err.Error())
				os.Exit(1)
			}
		default:
			fmt.Println("Invalid command")
			fmt.Println()
		}
	}
}

func commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Display this help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println()

	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)

	return nil
}
