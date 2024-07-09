package commands

import (
	"abhiroop43/pokedex/dto"
	"fmt"
)

func CommandPokedex(config *dto.Config, args ...string) error {

	fmt.Println("Pokemons currently in the Pokedex:")

	for _, pokemon := range config.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
