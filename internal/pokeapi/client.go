package pokeapi

import (
	"net/http"
	"time"

	"github.com/thealmightychuck/pokedexcli/internal/pokecache"
)

// create a client in main as it lives for the lifetime of the whole program.
// main build the client and needs a way to get it to objects or functions that process requests.
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(1 * time.Minute),
	}
}
