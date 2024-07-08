package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseUrl + endpoint

	// check cache entry
	cacheData, ok := c.cache.Get(fullUrl)
	if ok {
		//fmt.Println("loading from cache....")
		response := Pokemon{}
		err := json.Unmarshal(cacheData, &response)
		if err != nil {
			return Pokemon{}, err
		}

		return response, nil
	}

	//fmt.Println("loading from http request....")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	response := Pokemon{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return response, nil
}
