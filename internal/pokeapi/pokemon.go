package pokeapi

import (
	"encoding/json"
	"errors"
	"io"

	"net/http"
)

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		}
	}
	Caught bool
}

func FetchPokemonData(pokemonName string) (*Pokemon, error) {
	url := POKEMON_API_URL + "/pokemon/" + pokemonName

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("failed to fetch location")
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("pokemon not found")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("api responded with code: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	// chance := rand.IntN(350) // Mewtwo is 340
	// // fmt.Println("Chance:", chance)

	var pokemon Pokemon
	json.Unmarshal(body, &pokemon)

	return &pokemon, nil
}
