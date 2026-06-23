// nonmutexexample.go — Day 18: Intentional data race on a map (NO mutex).
// A writer goroutine and a reader goroutine access the same map concurrently
// without any synchronisation. This will panic or produce undefined behaviour.
// Run with `go run -race` to confirm the race. This is the "wrong" counterpart
// to mutexexample.go.
package main

import (
	"time"
)

func main() {
	// NO MUTEX - WRONG!
	var data map[string]int

	// Writer
	go func() {
		data["x"] = 1 // Modifying map structure
	}()

	// Reader
	go func() {
		_ = data["x"] // Simultaneously reading
	}()

	time.Sleep(time.Second * 1)
}
