package pokeapi

import (
	"net/http"
	"time"

	"github.com/samoei/pokedexcli/internal/pokeapi/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(expiryInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(expiryInterval),
	}
}
