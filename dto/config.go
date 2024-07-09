package dto

type Config struct {
	PokeApiClient       Client
	NextLocationUrl     *string
	PreviousLocationUrl *string
	CaughtPokemon       map[string]Pokemon
}
