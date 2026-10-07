package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

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

func MakeRequest(url string) (Response, error) {
	// create new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, err
	}

	// make request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	// reading response body into data variable
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	// unmarshal json data into a resonse struct
	var mapResponse Response
	if err := json.Unmarshal(data, &mapResponse); err != nil {
		return Response{}, err
	}

	return mapResponse, nil
}
