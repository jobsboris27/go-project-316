package checker

import (
	"sync"
)

type LinkCache struct {
	mu    sync.Mutex
	cache map[string]BrokenLink
}

func NewLinkCache() *LinkCache {
	return &LinkCache{
		cache: make(map[string]BrokenLink),
	}
}

func (c *LinkCache) Get(key string) (BrokenLink, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	result, ok := c.cache[key]
	return result, ok
}

func (c *LinkCache) Set(key string, value BrokenLink) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *LinkCache) Contains(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	return ok
}
