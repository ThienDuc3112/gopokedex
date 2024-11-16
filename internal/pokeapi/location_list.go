package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocation(pageURL *string) (RespLocationShallow, error) {
	getUrl := baseURL + "/location-area"
	if pageURL != nil {
		getUrl = *pageURL
	}

	if cached, ok := c.cache.Get(getUrl); ok {
		var locations RespLocationShallow
		if err := json.Unmarshal(cached, &locations); err != nil {
			return RespLocationShallow{}, fmt.Errorf("error decoding data: %v", err)
		}
		return locations, nil
	}

	resp, err := c.httpClient.Get(getUrl)
	if err != nil {
		return RespLocationShallow{}, fmt.Errorf("get location list failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationShallow{}, fmt.Errorf("unable to read request body: %v", err)
	}

	var locations RespLocationShallow
	if err = json.Unmarshal(body, &locations); err != nil {
		return RespLocationShallow{}, fmt.Errorf("error decoding data: %v", err)
	}

	c.cache.Add(getUrl, body)

	return locations, nil

}
