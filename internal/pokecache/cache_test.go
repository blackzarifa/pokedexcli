package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)

	// Test adding and retrieving
	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	// Check if we can get it back
	retrieved, found := cache.Get(key)
	if !found {
		t.Errorf("Expected to find key %s in cache", key)
	}

	if string(retrieved) != string(val) {
		t.Errorf("Expected %s, got %s", string(val), string(retrieved))
	}

	// Test reaping
	time.Sleep(150 * time.Millisecond)

	// The entry should be gone now
	_, found = cache.Get(key)
	if found {
		t.Errorf("Expected key %s to be reaped from cache", key)
	}
}
