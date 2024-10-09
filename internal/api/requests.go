package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
}

func GetLocations(url string) ([]Location, string, string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error making get request: %v\n", err)
		return []Location{}, "", "", err
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return []Location{}, "", "", err
	}
	if response.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
		return []Location{}, "", "", err
	}

	// Only define used fields from response
	type locationResponse struct {
		Next     string     `json:"next"`
		Previous string     `json:"previous"`
		Results  []Location `json:"results"`
	}

	locations := locationResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Printf("Can't decode response: %v", err)
		return []Location{}, "", "", err
	}

	return locations.Results, locations.Next, locations.Previous, nil
}
