// assignment.go — Day 16 Assignment: Channel pipeline pattern.
// generate sends 1..10 into a channel, square reads from that channel and
// sends squared values into another. Main ranges over the final channel.
// Each stage runs as its own goroutine, and defer close() propagates
// completion through the pipeline automatically.
package main

import (
	"fmt"
)

func generate(out chan<- int) {
	defer close(out)
	for i := 1; i <= 10; i++ {
		out <- i
	}
}

func square(in <-chan int, out chan<- int) {
	defer close(out)
	for numb := range in {
		out <- numb * numb

	}

}

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go generate(nums)
	go square(nums, squares)

	for v := range squares {
		fmt.Println(v)
	}
}
