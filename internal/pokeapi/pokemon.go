package pokeapi

import "encoding/json"

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height "`
	Weight         int    `json:"weight"`
	Stat           []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Pokedex struct {
	Dex map[string]Pokemon
}

func NewPokedex() Pokedex {
	p := Pokedex{
		Dex: map[string]Pokemon{},
	}
	return p
}

func CatchUnmarshal(data []byte) (Pokemon, error) {
	var catchResponse Pokemon
	if err := json.Unmarshal(data, &catchResponse); err != nil {
		return Pokemon{}, err
	}
	return catchResponse, nil
}
