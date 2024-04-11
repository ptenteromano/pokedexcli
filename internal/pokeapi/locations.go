package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []locationArea
}

func GetLocations(url string) LocationResponse {
	if url == "" {
		url = POKEMON_API_URL + "/location-area/"
	}

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the body of the response
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	// To print the raw JSON as a string, you can simply do:
	// fmt.Println("Raw JSON string:")

	var data LocationResponse
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	// fmt.Println(data.Results)

	return data
}
