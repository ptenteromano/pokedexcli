package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

	var data interface{}
	json.Unmarshal(body, &data)

	fmt.Println(data)
	return PokemonList{pokemon: []string{"bulbasaur", "charmander", "squirtle"}}
}
