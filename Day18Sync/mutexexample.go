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

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++

}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter

}

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
