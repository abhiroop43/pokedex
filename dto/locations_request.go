package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationsResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// check cache entry
	cacheData, ok := c.cache.Get(fullUrl)
	if ok {
		//fmt.Println("loading from cache....")
		locationsResponse := LocationsResponse{}
		err := json.Unmarshal(cacheData, &locationsResponse)
		if err != nil {
			return LocationsResponse{}, err
		}

		return locationsResponse, nil
	}

	//fmt.Println("loading from http request....")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode > 399 {
		return LocationsResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	locationsResponse := LocationsResponse{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return LocationsResponse{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationsResponse, nil
}
