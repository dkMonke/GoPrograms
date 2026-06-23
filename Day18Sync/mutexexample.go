// mutexexample.go — Day 18: sync.Mutex for safe concurrent access.
// Counter wraps an int with a Mutex. Inc and Value lock/unlock the mutex
// to prevent data races. Multiple goroutines safely increment and read
// the counter. Compare with nonmutexexample.go to see the difference.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter is a goroutine-safe integer counter. It pairs an int value with a
// sync.Mutex so that concurrent increments and reads are serialised and free
// of data races. The zero value (Counter{}) is ready to use; the mutex does
// not need explicit initialisation.
type Counter struct {
	mu      sync.Mutex
	counter int
}

// Inc atomically increments the counter by one. It acquires the mutex before
// touching the shared field and releases it via defer when the method returns,
// guaranteeing that no two goroutines mutate counter at the same time. It takes
// no arguments and returns nothing.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++

}

// Value returns the current counter value. It locks the mutex for the duration
// of the read (released via defer) so that the returned int reflects a
// consistent snapshot that cannot be torn by a concurrent Inc. It takes no
// arguments and returns the counter as an int.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter

}

// main demonstrates the mutex-protected Counter under concurrent access. It
// creates a Counter, spawns several goroutines that call Inc and Value, mixes
// in direct calls from the main goroutine, then sleeps 100ms to let the
// goroutines finish before printing the final value. The brief sleep is a
// teaching shortcut for goroutine synchronisation, not a robust wait.
func main() {
	c := Counter{}
	go func() {
	}()
	c.Inc()
	go func() {
		c.Inc()
	}()
	go func() {
		fmt.Println(c.Value())
	}()
	c.Inc()
	c.Inc()
	time.Sleep(time.Millisecond * 100)
	fmt.Println(c.Value())
}
