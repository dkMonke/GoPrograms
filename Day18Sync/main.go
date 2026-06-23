// Package main demonstrates a concurrency-safe in-memory cache guarded by a
// sync.RWMutex. It shows how multiple goroutines can write to and read from a
// shared map without data races, coordinated by a sync.WaitGroup.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache is a concurrency-safe key/value store mapping strings to strings.
// All access to the underlying map is synchronized through mu, allowing
// concurrent readers or a single exclusive writer at any given time.
type Cache struct {
	mu sync.RWMutex
	m  map[string]string
}

// NewCache creates and returns a pointer to a ready-to-use Cache with its
// internal map initialized. The returned Cache is safe for concurrent use.
func NewCache() *Cache {
	return &Cache{m: make(map[string]string)}
}

// Get retrieves the value stored under key k. It acquires a read lock so that
// multiple Get calls may proceed concurrently. It returns the value and true
// if the key is present, or the empty string and false if it is not.
func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}

// Set stores the value v under key k. It acquires an exclusive write lock to
// prevent concurrent reads or writes while the map is being modified. An
// existing value for the same key is overwritten.
func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

// main builds a Cache and exercises it concurrently. It launches ten writer
// goroutines that each Set a distinct "key%d"/"value%d" pair, then ten reader
// goroutines that each briefly sleep before Getting their corresponding key
// and printing the index, value, and presence flag. A sync.WaitGroup ensures
// main blocks until all twenty goroutines have completed. Because the readers
// and writers run concurrently, the output order is non-deterministic and a
// given key may or may not have been written by the time it is read.
func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			v, ok := cache.Get(fmt.Sprintf("key%d", i))
			fmt.Println(i, v, ok)
		}(i)
	}
	wg.Wait()
}
