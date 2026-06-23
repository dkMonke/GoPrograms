// main.go — Day 15: Goroutines and sync.WaitGroup.
// Launches 5 concurrent worker goroutines. Each worker sleeps proportionally
// to its ID. sync.WaitGroup ensures main waits for all workers to finish
// before printing "all done". defer wg.Done() guarantees the counter decrements.
package main

import (
	"fmt"
	"sync"
	"time"
)

// work simulates a unit of concurrent work for the worker identified by id.
// It sleeps for a duration proportional to id (id * 100ms) to mimic
// variable-length tasks, printing messages when it starts and finishes. The
// wg pointer lets it signal completion to the caller; defer wg.Done()
// guarantees the WaitGroup counter is decremented even if the function
// returns early.
func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("worker %d dones\n", id)
}

// main launches five work goroutines and uses a sync.WaitGroup to wait for
// all of them to complete. It calls wg.Add(1) before starting each goroutine
// and wg.Wait() blocks until every worker has called wg.Done(), after which
// "all done" is printed. This ensures main does not exit prematurely.
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done")
}
