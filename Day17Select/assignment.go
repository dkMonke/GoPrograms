// assignment.go — Day 17 Assignment: Fan-in concurrency pattern.
// fanIn merges multiple input channels into a single output channel.
// A goroutine is spawned per input channel to forward values. sync.WaitGroup
// tracks when all inputs are exhausted, then closes the output channel.
// Main creates three producers with different rates and merges them.
package main

import (
	"fmt"
	"sync"
	"time"
)

func fanIn(inputs []<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(100 * time.Millisecond)
			ch1 <- fmt.Sprintf("ch1-%d", i)
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(150 * time.Millisecond)
			ch2 <- fmt.Sprintf("ch2-%d", i)
		}
		close(ch2)
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(300 * time.Millisecond)
			ch3 <- fmt.Sprintf("ch3-%d", i)
		}
		close(ch3)
	}()

	merged := fanIn([]<-chan string{ch1, ch2, ch3})

	for val := range merged {
		fmt.Println(val)
	}
}
