package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(nameOrId string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + nameOrId

	data, ok := c.cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		data, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(fullURL, data)
	}
	result := Pokemon{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
