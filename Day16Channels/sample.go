// sample.go — Day 16: Simple producer-consumer with an unbuffered channel.
// producer sends 0..4 into ch with a 2-second delay between sends, then closes
// the channel. Main uses range to receive values until the channel is closed,
// demonstrating blocking send/receive semantics of unbuffered channels.
package main

import (
	"fmt"
	"time"
)

// producer sends the integers 0 through 4 into the send-only channel ch,
// pausing two seconds between sends to simulate slow production. It closes ch
// via defer once all values are sent so a ranging consumer knows to stop.
// Because ch is unbuffered, each send blocks until the consumer receives.
func producer(ch chan<- int) {

	defer close(ch)

	for i := 0; i < 5; i++ {

		ch <- i
		time.Sleep(time.Second * 2)

	}

}

// main creates an unbuffered channel, runs producer as a goroutine to feed it,
// and ranges over the channel to receive and print each value as it arrives.
// The range loop blocks waiting for each send and exits when producer closes
// the channel.
func main() {

	ch := make(chan int)

	go producer(ch)
	for v := range ch {
		fmt.Println(v)
	}

}
