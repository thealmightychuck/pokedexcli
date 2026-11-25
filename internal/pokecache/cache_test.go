package pokecache

import (
	"bytes"
	"testing"
	"time"
	//"github.com/thealmightychuck/pokedexcli/internal/pokecache"
)

func TestAddGet(t *testing.T) {

	const interval = 5 * time.Second
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	byteData, exists := cache.Get("https://example.com")

	expected := []byte("testdata")
	if !bytes.Equal(expected, byteData) {
		t.Errorf("cache.Get() expected data %v got %v", expected, byteData)
	}
	if !exists {
		t.Errorf("cache.Get() expected bool %v got %v", true, exists)
	}

	// Missing Key
	missingByteData, missingExists := cache.Get("https://testdata.com")
	if !bytes.Equal(nil, missingByteData) {
		t.Errorf("cache.Get() missing data, expected %v, got %v", nil, missingByteData)
	}
	if missingExists {
		t.Errorf("cache.Get() missing bool, expected %v, got %v", false, missingExists)
	}

	// Add replace the existing value for a same key.
	cache.Add("https://helloworld.com", []byte("helloworld"))
	cache.Add("https://helloworld.com", []byte("helloworld2"))
	byteData2, exists2 := cache.Get("https://helloworld.com")

	expected2 := []byte("helloworld2")
	if !bytes.Equal(expected2, byteData2) {
		t.Errorf("cache.Get() expected data %v got %v", expected2, byteData2)
	}
	if !exists2 {
		t.Errorf("cache.Get() expected bool %v got %v", true, exists2)
	}

	// Multiple different keys with the same cache
	cache.Add("https://EntryA.com", []byte("entrya"))
	cache.Add("https://EntryB.com", []byte("entryb"))

	cases := []struct {
		key      string
		expected []byte
	}{
		{"https://EntryA.com", []byte("entrya")},
		{"https://EntryB.com", []byte("entryb")},
	}

	for _, c := range cases {
		got, exists := cache.Get(c.key)
		if !bytes.Equal(c.expected, got) {
			t.Errorf("cache.Get(%q) expected data %v got %v", c.key, c.expected, got)
		}
		if !exists {
			t.Errorf("cache.Get(%q) expected bool true got %v", c.key, exists)
		}
	}

	// adding tests checking that cache is deleted correctly.
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(20 * time.Millisecond)
	cache.Add("page", []byte("locations"))
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := cache.Get("page"); !ok {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("expected the background loop to remove the expired entry")
}
