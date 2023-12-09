package pokecache

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

	go c.loopRemove(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) Remove(interval time.Duration) {
	tiemBefore := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(tiemBefore) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) loopRemove(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Remove(interval)
	}
}
