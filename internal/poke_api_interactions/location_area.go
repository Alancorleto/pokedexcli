package pokeapiinteractions

import (
	"fmt"
)

type LocationsResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var nextURL string
var previousURL string

func Map() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if nextURL != "" {
		url = nextURL
	}

	getAndPrintLocations(url)
	return nil
}

func Mapb() error {
	if previousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := previousURL

	getAndPrintLocations(url)
	return nil
}

func getAndPrintLocations(url string) error {
	locations, err := Get[LocationsResponse](url)
	if err != nil {
		return err
	}

	for i := range len(locations.Results) {
		fmt.Printf("%s\n", locations.Results[i].Name)
	}

	nextURL = locations.Next
	previousURL = locations.Previous

	return nil
}
