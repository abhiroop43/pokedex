package cache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.DeleteContinuously(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	retrievedCache, ok := c.cache[key]

	return retrievedCache.val, ok
}

func (c *Cache) DeleteContinuously(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.Delete(interval)
	}
}

func (c *Cache) Delete(t time.Duration) {
	threshold := time.Now().UTC().Add(-t)
	for k, v := range c.cache {
		if v.createdAt.Before(threshold) {
			delete(c.cache, k)
		}
	}
}
