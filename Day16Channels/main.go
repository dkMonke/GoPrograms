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

// worker consumes jobs from the receive-only jobs channel until it is closed.
// For each job it prints a processing message, sleeps a random number of
// seconds (0-9) to simulate variable work, and sends the doubled job value to
// the send-only results channel. id identifies the worker in log output.
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Printf("worker %d processing %d\n", id, j)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		results <- j * 2
	}

}

// main implements a worker-pool pattern. It creates buffered jobs and results
// channels, starts three worker goroutines, enqueues five jobs, and closes the
// jobs channel to signal the workers to stop once the queue is drained. It
// then receives and prints all five results. Buffering lets the producer
// enqueue without blocking on slow workers.
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
