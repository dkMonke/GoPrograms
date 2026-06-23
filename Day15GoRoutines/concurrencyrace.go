// concurrencyrace.go — Day 15: Data race demonstration.
// 101 goroutines increment a shared counter without synchronisation.
// This is intentionally buggy — counter++ is not atomic, so the final value
// is unpredictable. Run with `go run -race` to detect the race condition.
package main

import (
	"fmt"
	"time"
)

// main launches 101 goroutines that each increment a shared counter variable
// with no synchronisation. Because counter++ is a read-modify-write that is
// not atomic, concurrent goroutines race on it and the printed total is
// non-deterministic, typically less than 101. The time.Sleep gives the
// goroutines a chance to run before the value is printed. Run with the
// `-race` flag to have the Go runtime detect and report the data race.
func main() {
	counter := 0

	for i := 0; i <= 100; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter)
}
