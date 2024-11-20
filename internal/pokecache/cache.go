package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex *sync.Mutex
	cache map[string]CacheEntry
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.cache[key]
	if !ok {
		return value.val, ok
	}
	return value.val, ok
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mutex: &sync.Mutex{},
		cache: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cache {
		if time.Since(v.createdAt) > interval {
			delete(c.cache, k)
		}
	}
}
