package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	interval     time.Duration
	mutex        sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.cacheEntries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for key, entry := range c.cacheEntries {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}
