package commands

import (
	"abhiroop43/pokedex/dto"
	"errors"
	"fmt"
)

func CommandMapB(config *dto.Config, args ...string) error {

	if config.PreviousLocationUrl == nil {
		return errors.New("you are on the first page")
	}

	locations, err := config.PokeApiClient.ListLocations(config.PreviousLocationUrl)
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
