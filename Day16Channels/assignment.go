// assignment.go — Day 16 Assignment: Channel pipeline pattern.
// generate sends 1..10 into a channel, square reads from that channel and
// sends squared values into another. Main ranges over the final channel.
// Each stage runs as its own goroutine, and defer close() propagates
// completion through the pipeline automatically.
package main

import (
	"fmt"
)

// generate is the first stage of the pipeline. It sends the integers 1
// through 10 into the send-only channel out and, via defer close(out), closes
// the channel once all values have been sent. Closing signals downstream
// stages that no more values will arrive so their range loops can terminate.
func generate(out chan<- int) {
	defer close(out)
	for i := 1; i <= 10; i++ {
		out <- i
	}
}

// square is the middle stage of the pipeline. It ranges over the receive-only
// channel in until that channel is closed, sends each value squared into the
// send-only channel out, and closes out (via defer) when the input is
// exhausted, propagating completion to the final consumer.
func square(in <-chan int, out chan<- int) {
	defer close(out)
	for numb := range in {
		out <- numb * numb

	}

}

// main wires up the channel pipeline: it creates the nums and squares
// channels, runs generate and square as goroutines connected through them,
// and ranges over the squares channel to print each squared value. The loop
// ends automatically when square closes the squares channel.
func main() {
	nums := make(chan int)
	squares := make(chan int)

	go generate(nums)
	go square(nums, squares)

	for v := range squares {
		fmt.Println(v)
	}
}
