// Package main demonstrates a simple concurrency-safe rate limiter.
//
// It implements a fixed-window rate limiter that caps the number of allowed
// operations per one-second window, then exercises it from many goroutines to
// observe how many requests are permitted over a short interval.
package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter is a fixed-window rate limiter that permits up to a configured
// number of operations within each one-second window.
//
// Fields:
//   - mu:        guards count and startTime for concurrent access.
//   - count:     number of operations allowed in the current window.
//   - startTime: the start time of the current window.
//   - perSecond: the maximum number of operations allowed per one-second window.
type RateLimiter struct {
	mu        sync.Mutex
	count     int
	startTime time.Time
	perSecond int
}

// NewRateLimiter creates and returns a new RateLimiter whose window begins at
// the current time.
//
// The perSecond parameter is intended to set the maximum number of operations
// allowed per one-second window. Note that the returned RateLimiter does not
// store perSecond (it is left at its zero value), so the limit is effectively
// zero unless the field is set separately.
func NewRateLimiter(perSecond int) *RateLimiter {
	return &RateLimiter{count: 0, startTime: time.Now()}
}

// Allow reports whether an operation is permitted under the current rate limit
// and, when permitted, records it against the current window.
//
// It resets the window (clearing count and resetting startTime) once one second
// has elapsed since the window began. If the number of operations already
// allowed in the window has reached perSecond, it returns false; otherwise it
// increments the count and returns true.
//
// Allow takes the limiter's mutex to make the check-and-update atomic across
// goroutines.
func (r *RateLimiter) Allow() bool {

	r.mu.Lock()
	defer r.mu.Lock()

	duration := time.Now().Sub(r.startTime) // Returns time.Duration
	seconds := duration.Seconds()
	if seconds >= 1.0 {
		r.count = 0
		r.startTime = time.Now()
	}
	if r.count >= r.perSecond {
		return false
	}
	r.count++
	return true
}

// main exercises the RateLimiter by launching 100 goroutines that repeatedly
// call Allow for two seconds.
//
// Each goroutine loops until two seconds have passed since start, calling Allow
// on every iteration and sleeping 10 milliseconds between attempts. Allowed
// operations increment a shared successCount (guarded by its own mutex) and are
// logged, while denied operations are also logged. After all goroutines finish,
// it prints the total number of allowed operations alongside the expected count.
func main() {

	rateLimiter := NewRateLimiter(2)
	var wg sync.WaitGroup
	var successCount int
	var mu sync.Mutex

	start := time.Now()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				if time.Since(start) > 2*time.Second {
					break
				}
				if rateLimiter.Allow() {
					mu.Lock()
					successCount++
					mu.Unlock()
					fmt.Printf("Goroutine %d: ALLOWED (total: %d)\n", id, successCount)

				} else {
					fmt.Printf("Goroutine %d: Denied\n", id)
				}
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nTotal allowed: %d\n", successCount)
	fmt.Printf("Expected: ~20 (10 per second x 2 seconds)\n")

}
