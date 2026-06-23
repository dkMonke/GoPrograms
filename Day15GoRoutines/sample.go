// sample.go — Day 15: Basic goroutine example.
// say("Hello") runs concurrently as a goroutine while say("world") runs on main.
// The "world" goroutine sleeps 100ms per iteration to give "Hello" time to print.
// Without the sleep, main might finish before the "Hello" goroutine completes.
package main

import (
	"fmt"
	"time"
)

// say prints the string s together with an index three times. When s equals
// "world" it sleeps 100ms after each print, deliberately yielding the
// processor so a concurrently running say goroutine has time to interleave
// its output. This demonstrates cooperative scheduling between goroutines.
func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, i)
		if s == "world" {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// main runs say("Hello") concurrently as a goroutine while running
// say("world") directly on the main goroutine. The "world" call blocks main
// long enough (via its sleeps) for the "Hello" goroutine to print before the
// program exits, illustrating basic goroutine concurrency.
func main() {
	go say("Hello")
	say("world")
}
