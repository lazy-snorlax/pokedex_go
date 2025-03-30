package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]CacheEntry
	mu           *sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: map[string]CacheEntry{},
		mu:           &sync.RWMutex{},
	}
	// Start the reapLoop in a separate goroutine for background cleanup
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = CacheEntry{createdAt: time.Now(), val: val}
	fmt.Printf("Adding to cache - key: %v  val: %v \n", key, val)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.cacheEntries[key]
	if exists {
		fmt.Printf("Fetching from cache - key: %v  val: %v \n", key, entry.val)
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop() // Ensure ticker stops to prevent resource leaks when reapLoop exits

	for {
		select {
		case <-ticker.C: // On every tick:
			c.mu.Lock() // Safely lock the cache
			for key, entry := range c.cacheEntries {
				if time.Since(entry.createdAt) > interval { // Check expiration
					delete(c.cacheEntries, key)
				}
			}
			c.mu.Unlock() // Unlock after cleaning
		}
	}
}
