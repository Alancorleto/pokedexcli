package pokeapiinteractions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func PrintLocations() {
	url := "https://pokeapi.co/api/v2/location/"
	data, err := MakeGetRequest(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	locations, err := DecodeResponseAsLocations(data)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	limit := 20
	if len(locations.Results) < limit {
		limit = len(locations.Results)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%s\n", locations.Results[i].Name)
	}
}

func MakeGetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP error: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DecodeResponseAsLocations(data []byte) (*LocationsResponse, error) {
	var locations LocationsResponse
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}
	return &locations, nil
}
