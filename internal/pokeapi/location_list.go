package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationsResp := RespShallowLocations{}

	cacheEntry, exists := c.cache.Get(url)
	fmt.Printf("Checking cache - URL: %v  CacheEntry: %v  Exists: %v \n", url, cacheEntry, exists)
	if exists {
		fmt.Println("Pulling from cache")
		err := json.Unmarshal(cacheEntry, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, nil
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, nil
	}
	c.cache.Add(url, dat)

	return locationsResp, nil
}
