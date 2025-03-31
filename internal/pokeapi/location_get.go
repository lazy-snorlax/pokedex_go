package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	// Check cache for result
	cacheEntry, exists := c.cache.Get(url)
	if exists {
		locationResp := Location{}
		fmt.Println("Pulling from cache")
		err := json.Unmarshal(cacheEntry, &locationResp)
		if err != nil {
			return Location{}, nil
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, nil
	}
	c.cache.Add(url, dat)

	return locationResp, nil
}
