package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	interval   time.Duration
	mu         sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]cacheEntry),
		interval:   interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntry[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntry, key)
			}
		}

		c.mu.Unlock()
	}
}
