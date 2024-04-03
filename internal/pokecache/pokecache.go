package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{cache: make(map[string]cacheEntry), mu: sync.Mutex{}}
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return val.val, true
}
