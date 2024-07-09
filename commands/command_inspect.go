package commands

import (
	"abhiroop43/pokedex/dto"
	"errors"
	"fmt"
)

func CommandInspect(config *dto.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]

	pokemonDetails, ok := config.CaughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemonDetails.Name)
	fmt.Printf("Height: %v\n", pokemonDetails.Height)
	fmt.Printf("Weight: %v\n", pokemonDetails.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemonDetails.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typ := range pokemonDetails.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}
