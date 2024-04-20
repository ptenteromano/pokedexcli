package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"
)

type catchResponse struct {
	BaseExperience int `json:"base_experience"`
}

func AttemptCatch(pokemonName string) (bool, error) {
	url := POKEMON_API_URL + "/pokemon/" + pokemonName

	resp, err := http.Get(url)

	if err != nil {
		return false, errors.New("failed to fetch location")
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return false, errors.New("pokemon not found")
	}

	if resp.StatusCode != 200 {
		return false, errors.New("api responded with code: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, errors.New("failed to read response body")
	}

	chance := rand.IntN(350) // Mewtwo is 340
	// fmt.Println("Chance:", chance)

	var data catchResponse
	json.Unmarshal(body, &data)
	// fmt.Println("Base experience:", data.BaseExperience)

	if chance < data.BaseExperience {
		return false, nil
	}

	return true, nil
}
