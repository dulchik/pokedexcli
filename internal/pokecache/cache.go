package pokecache

import (
	"time"
	"sync"
)

var mu sync.Mutex

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type cache struct {
	map[string]cacheEntry
}

func (c cache) NewCache(interval time.Duration) error {
	mu.Lock()
	defer mu.Unlock()
	return nil
}

func (c cache) Add(key string, val []byte) {
	
}

func (c cache) Get(key string) []byte bool{
	
}

func (c cache) reapLoop() {

}
