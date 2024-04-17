package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type exploreResponse struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type PokemonList struct {
	pokemon []string
}

func ExploreLocation(locationName string) PokemonList {
	url := POKEMON_API_URL + "/location-area/" + locationName

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	var data exploreResponse
	json.Unmarshal(body, &data)

	// Pretty print json
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	log.Println(string(jsonData))
	return PokemonList{pokemon: []string{"bulbasaur", "charmander", "squirtle"}}
}
