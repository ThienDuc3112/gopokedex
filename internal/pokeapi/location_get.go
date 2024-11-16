package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocation(name string) (Location, error) {
	URL := baseURL + "/location-area/" + name
	if data, ok := c.cache.Get(URL); ok {
		var location Location
		if err := json.Unmarshal(data, &location); err != nil {
			return Location{}, fmt.Errorf("error decoding data: %v", err)
		}
		return location, nil
	}

	resp, err := c.httpClient.Get(URL)
	if err != nil {
		return Location{}, fmt.Errorf("get location failed: %v", err)
	}
	if resp.StatusCode > 299 {
		return Location{}, fmt.Errorf("location not found: %v", name)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, fmt.Errorf("unable to read request body: %v", err)
	}

	var location Location
	if err := json.Unmarshal(data, &location); err != nil {
		return Location{}, fmt.Errorf("error decoding data: %v", err)
	}

	c.cache.Add(URL, data)
	return location, nil
}
