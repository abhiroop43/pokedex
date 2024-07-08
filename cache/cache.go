package cache

import (
	"time"
)

type Cache struct {
	Cache map[string]CacheEntry
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func (c *Cache) Add(key string, val []byte) {
	c.Cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	retrievedCache, ok := c.Cache[key]

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
	for k, v := range c.Cache {
		if v.createdAt.Before(threshold) {
			delete(c.Cache, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cache: make(map[string]CacheEntry),
	}
	go c.DeleteContinuously(interval)
	return c
}
