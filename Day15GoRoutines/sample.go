// sample.go — Day 15: Basic goroutine example.
// say("Hello") runs concurrently as a goroutine while say("world") runs on main.
// The "world" goroutine sleeps 100ms per iteration to give "Hello" time to print.
// Without the sleep, main might finish before the "Hello" goroutine completes.
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, i)
		if s == "world" {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	go say("Hello")
	say("world")
}
