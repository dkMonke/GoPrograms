// nonmutexexample.go — Day 18: Intentional data race on a map (NO mutex).
// A writer goroutine and a reader goroutine access the same map concurrently
// without any synchronisation. This will panic or produce undefined behaviour.
// Run with `go run -race` to confirm the race. This is the "wrong" counterpart
// to mutexexample.go.
package main

import (
	"time"
)

// main intentionally triggers a data race to illustrate why synchronisation is
// required. It launches one goroutine that writes to a shared map and another
// that reads from it concurrently, with no mutex or other coordination. Because
// the map is also nil (never made), the write can panic; under `go run -race`
// the race detector reports concurrent access. It takes no arguments and the
// 1-second sleep merely keeps main alive long enough for the goroutines to run.
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
