package pokeapi

import "encoding/json"

type ExplorationEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func ExploreUnmarshal(data []byte) (ExplorationEncounters, error) {
	var exploreResponse ExplorationEncounters
	if err := json.Unmarshal(data, &exploreResponse); err != nil {
		return ExplorationEncounters{}, err
	}
	return exploreResponse, nil
}
