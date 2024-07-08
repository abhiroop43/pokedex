package dto

import (
	"abhiroop43/pokedex/cache"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
