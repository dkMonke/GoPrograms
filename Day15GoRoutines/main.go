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

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("worker %d dones\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done")
}
