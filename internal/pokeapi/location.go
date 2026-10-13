package pokeapi

import "encoding/json"

type Response struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func MapUnmarshal(data []byte) (Response, error) {
	// unmarshal json data into a resonse struct
	var mapResponse Response
	if err := json.Unmarshal(data, &mapResponse); err != nil {
		return Response{}, err
	}

	return mapResponse, nil
}
