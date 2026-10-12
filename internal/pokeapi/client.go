package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ashuffman/pokedexcli/internal/pokecache"
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

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	c := Client{
		client: http.Client{},
		cache:  pokecache.NewCache(interval),
	}
	return c
}

func (c Client) MakeRequest(url string) (Response, error) {
	// check the cache
	data, ok := c.cache.Get(url)
	if !ok {
		// create new request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Response{}, err
		}

		// make request
		res, err := c.client.Do(req)
		if err != nil {
			return Response{}, err
		}
		defer res.Body.Close()

		// reading response body into data variable
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Response{}, err
		}

		// cache the response
		c.cache.Add(url, data)
	}

	// unmarshal json data into a resonse struct
	var mapResponse Response
	if err := json.Unmarshal(data, &mapResponse); err != nil {
		return Response{}, err
	}

	return mapResponse, nil
}
