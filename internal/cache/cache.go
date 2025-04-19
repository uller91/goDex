package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Data	map[string]cacheEntry
	mu		*sync.Mutex
}

type cacheEntry struct {
	CreatedAt	time.Time
	Val			[]byte
}

func NewCache(interval time.Duration) *Cache {	//interval is needed to show the duration of life of the Cache entry
	cache := &Cache{
		Data: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval) //here

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if c.Data == nil {
		c.Data = make(map[string]cacheEntry)
	}

	c.Data[key] = cacheEntry{
        CreatedAt: time.Now(),
        Val:       val,
    }
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.Data[key]
	if !ok {
		return nil, false
	}
	
	return value.Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		for key, entry := range c.Data {
			age := time.Since(entry.CreatedAt)
			if age > interval {
				delete(c.Data, key)
			}
		}

		c.mu.Unlock()
	}
}