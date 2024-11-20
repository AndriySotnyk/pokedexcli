package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(nameOrId string) (Location, error) {
	fullURL := baseURL + "/location-area/" + nameOrId

	data, ok := c.cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return Location{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}

		data, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			return Location{}, err
		}
		c.cache.Add(fullURL, data)
	}
	result := Location{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
