package main

import "abhiroop43/pokedex/dto"

func main() {
	config := dto.Config{
		PokeApiClient: dto.Client{},
	}
	startRepl(&config)
}
