package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	MU      sync.Mutex
	Entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		MU:      sync.Mutex{},
		Entries: map[string]cacheEntry{},
	}
	c := &newCache
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.MU.Lock()
	defer c.MU.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.Entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.MU.Lock()
	defer c.MU.Unlock()
	entry, ok := c.Entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.MU.Lock()
		for key := range c.Entries {
			age := time.Since(c.Entries[key].createdAt)
			if age > interval {
				delete(c.Entries, key)
			}
		}
		c.MU.Unlock()
	}
}
