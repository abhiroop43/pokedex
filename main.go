package main

import (
	"abhiroop43/pokedex/dto"
	"time"
)

func main() {
	config := dto.Config{
		PokeApiClient: dto.NewClient(time.Hour),
		CaughtPokemon: make(map[string]dto.Pokemon),
	}
	startRepl(&config)
}
