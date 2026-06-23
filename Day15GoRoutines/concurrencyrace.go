// concurrencyrace.go — Day 15: Data race demonstration.
// 101 goroutines increment a shared counter without synchronisation.
// This is intentionally buggy — counter++ is not atomic, so the final value
// is unpredictable. Run with `go run -race` to detect the race condition.
package main

import (
	"fmt"
	"time"
)

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
