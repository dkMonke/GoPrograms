// sample.go — Day 16: Simple producer-consumer with an unbuffered channel.
// producer sends 0..4 into ch with a 2-second delay between sends, then closes
// the channel. Main uses range to receive values until the channel is closed,
// demonstrating blocking send/receive semantics of unbuffered channels.
package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {

	defer close(ch)

	for i := 0; i < 5; i++ {

		ch <- i
		time.Sleep(time.Second * 2)

	}

}

func main() {

	ch := make(chan int)

	go producer(ch)
	for v := range ch {
		fmt.Println(v)
	}

}
