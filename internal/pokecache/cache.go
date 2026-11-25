package pokecache

import (
	"sync"
	"time"
)

// Key is the URL form pokeapi.
// value is raw JSON bytes. json.Unmarshal will read data and convert it into a go
// data from map for cached data
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cached memory
type Cache struct {
	data     map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

// NewCache is a function that builds a brand new Cache with an initialized (but empty) map,
// and gives back a pointer to it so all future interactions share the same data.
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	// each time an interval is passed to NewCache, it should remove entries older than interval.

	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.data {
			if time.Since(value.createdAt) >= c.interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}

	//return
}
