package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []result `json:"results"`
}

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(*config) error {
	// create new request
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("error creating request: ", err)
		return err
	}

	// make request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error making request: ", err)
		return err
	}
	defer res.Body.Close()

	// reading response body into data variable
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response struct{}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	return nil
}
