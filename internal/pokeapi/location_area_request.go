package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	//check cache

	dat, ok := c.cache.Get(fullURL)

	if ok {
		//cache hit
		fmt.Println("--->cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)

		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss!<---")
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(areaName string) (LocationAreaResp, error) {
	endpoint := "/location-area/" + areaName
	fullURL := baseURL + endpoint
	//check cache

	dat, ok := c.cache.Get(fullURL)

	if ok {
		//cache hit

		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locationAreaResp)

		if err != nil {
			return LocationAreaResp{}, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationAreaResp)

	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationAreaResp, nil
}
