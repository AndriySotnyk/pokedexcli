package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	fullURL := baseURL + "/location-area"

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)

	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)

		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)

		if err != nil {
			return RespShallowLocations{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)

		if err != nil {
			return RespShallowLocations{}, err
		}
		c.cache.Add(fullURL, data)
		fmt.Println("cached")
	}

	locationResp := RespShallowLocations{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return locationResp, err
	}
	return locationResp, nil
}
