package commands

import (
	"abhiroop43/pokedex/dto"
	"errors"
	"fmt"
	"math/rand"
)

func CommandCatch(config *dto.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]

	pokemonDetails, err := config.PokeApiClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemonDetails.BaseExperience)

	if randNum >= threshold {
		return fmt.Errorf("failed to catch: %s", pokemonName)
	}

	config.CaughtPokemon[pokemonName] = pokemonDetails

	fmt.Printf("Pokemon %s was caught!\n", pokemonName)

	return nil
}
