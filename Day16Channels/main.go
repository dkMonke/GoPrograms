// main.go — Day 16: Worker pool pattern using channels.
// 3 workers read jobs from a buffered channel and write results to another.
// jobs is a receive-only channel (<-chan) for workers; results is send-only (chan<-).
// Closing the jobs channel signals workers to stop. Main collects all 5 results.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Printf("worker %d processing %d\n", id, j)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		results <- j * 2
	}

}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		fmt.Println("result:", <-results)
	}
}
