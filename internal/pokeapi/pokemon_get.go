package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	pageURL := baseURL + "/pokemon/" + name

	if data, ok := c.cache.Get(pageURL); ok {
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("cannot decode data: %v", err)
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Get(pageURL)
	if err != nil {
		return Pokemon{}, fmt.Errorf("cannot fetch data: %v", err)
	}
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("pokemon not found: %v", name)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("unable to read request body: %v", err)
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("cannot decode data: %v", err)
	}

	c.cache.Add(pageURL, body)

	return pokemon, nil
}
