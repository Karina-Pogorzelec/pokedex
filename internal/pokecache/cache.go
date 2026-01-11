package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val		  []byte	
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu   	 sync.Mutex
	interval time.Duration
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce := cacheEntry{
		createdAt:	time.Now(),
		val:		val,
	}

	c.cacheMap[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheMap[key]
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
		
		now := time.Now()
		for key, entry := range c.cacheMap {
			timeFromCreate := now.Sub(entry.createdAt)
			if timeFromCreate > c.interval {
				delete(c.cacheMap, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap:	make(map[string]cacheEntry),
		interval:	interval,
	}

	go c.reapLoop()

	return c
}