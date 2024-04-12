package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]CacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Data: make(map[string]CacheEntry),
		mu:   &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
