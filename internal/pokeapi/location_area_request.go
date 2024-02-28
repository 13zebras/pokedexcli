package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//////////////////////////////////////
// ListLocationAreas
//////////////////////////////////////

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	// location-area?offset=0&limit=20 because first page is actually offset=0
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	// fmt.Println("fullURL:", fullURL)
	// check cache for fullURL
	// if found, return cached data
	data, ok := c.cache.Get(fullURL)
	if ok {
		// fmt.Println("cache hit")
		responseLocationAreas := LocationAreasResponse{}
		err := json.Unmarshal(data, &responseLocationAreas)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return responseLocationAreas, nil
	}

	// fmt.Println("no cached values")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	responseLocationAreas := LocationAreasResponse{}
	err = json.Unmarshal(data, &responseLocationAreas)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	// Add data to cache before returning
	c.cache.Add(fullURL, data)

	return responseLocationAreas, nil

}

// ////////////////////////////////////
// GetLocationArea
// ////////////////////////////////////
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName

	fullURL := baseURL + endpoint

	// fmt.Println("fullURL:", fullURL)
	// check cache for fullURL
	// if found, return cached data
	data, ok := c.cache.Get(fullURL)
	if ok {
		// fmt.Println("cache hit")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	// fmt.Println("no cached values")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	// Add data to cache before returning
	c.cache.Add(fullURL, data)

	return locationArea, nil

}
