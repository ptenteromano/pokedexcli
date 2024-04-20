package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func ExploreLocation(locationName string) ([]string, error) {
	url := POKEMON_API_URL + "/location-area/" + locationName

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("failed to fetch location")
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("location not found")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("api responded with code: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	var data exploreResponse
	json.Unmarshal(body, &data)
	fmt.Println(data.getPokemon())

	return data.getPokemon(), nil
}

func (e exploreResponse) getPokemon() []string {
	pokemon := []string{}
	for _, encounter := range e.PokemonEncounters {
		pokemon = append(pokemon, encounter.Pokemon.Name)
	}

	return pokemon
}
