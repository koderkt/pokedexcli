package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {

	cache := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}
	go cache.readLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]

	return val.val, ok
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reader(interval)
	}
}

func (c *Cache) reader(interval time.Duration) {
	// Logic to delete cache entry once it expires
	c.mu.Lock()
	defer c.mu.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)

	for key, val := range c.cache {
		if val.createAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}
