package commands

import (
	"abhiroop43/pokedex/dto"
	"fmt"
)

func CommandMap() error {
	pokeApiClient := dto.Client{}
	locations, err := pokeApiClient.ListLocations()
	if err != nil {
		_ = fmt.Errorf("error fetching locations: %v", err)
	}
	for _, locations := range locations.Results {
		fmt.Printf("Location: %v, URL: %v\n", locations.Name, locations.URL)
	}
	return nil
}
