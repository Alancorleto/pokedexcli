package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mutex   sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

var cache *Cache

func NewCache(interval time.Duration) *Cache {
	cache = &Cache{
		entries: make(map[string]CacheEntry),
		mutex:   sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func Get(key string) ([]byte, bool) {
	return cache.Get(key)
}

func Add(key string, val []byte) {
	cache.Add(key, val)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	for range time.NewTicker(interval).C {
		c.mutex.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}
