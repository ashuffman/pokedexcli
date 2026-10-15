package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ashuffman/pokedexcli/internal/pokecache"
)

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

func (c Client) MakeRequest(url string) ([]byte, error) {
	// check the cache
	data, ok := c.cache.Get(url)
	if !ok {
		// create new request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []byte{}, err
		}

		// make request
		res, err := c.client.Do(req)
		if err != nil {
			return []byte{}, err
		}
		defer res.Body.Close()

		// check for status code, and return an error on an unsuccuessful call
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("bad response status: %s", res.Status)
		}

		// reading response body into data variable
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, err
		}

		// cache the response
		c.cache.Add(url, data)
	}

	return data, nil
}
