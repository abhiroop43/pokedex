package commands

import (
	"abhiroop43/pokedex/dto"
	"fmt"
)

func CommandMap(config *dto.Config, args ...string) error {
	locations, err := config.PokeApiClient.ListLocations(config.NextLocationUrl)
	if err != nil {
		_ = fmt.Errorf("error fetching locations: %v", err)
	}
	for _, locations := range locations.Results {
		fmt.Printf(" - Location: %v, URL: %v\n", locations.Name, locations.URL)
	}
	config.NextLocationUrl = locations.Next
	config.PreviousLocationUrl = locations.Previous
	return nil
}
