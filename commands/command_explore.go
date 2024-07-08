package commands

import (
	"abhiroop43/pokedex/dto"
	"errors"
	"fmt"
)

func CommandExplore(config *dto.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location provided")
	}

	locationName := args[0]

	//fmt.Printf("Location selected: %s\n\n", locationName)

	locationDetails, err := config.PokeApiClient.GetLocationDetails(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemons in %s:\n\n", locationDetails.Name)

	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
