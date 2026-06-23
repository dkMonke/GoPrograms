// main.go — Day 17: Select statement for timeout and cancellation.
// slowOperation uses select to wait on either a 3-second timer (simulating work)
// or a done channel (cancellation signal). The main goroutine closes done after
// 8 seconds. Since 3 < 8, the operation completes successfully here.
package main

import (
	"fmt"
	"time"
)

// slowOperation simulates a long-running task using a select statement that
// races two channels: a 3-second timer (time.After) representing the work
// completing, and the done cancellation channel. If the timer fires first it
// returns ("result", nil); if done is closed or receives first it returns an
// empty string and a "cancelled" error. The branch that becomes ready first
// wins.
func slowOperation(done <-chan struct{}) (string, error) {
	select {
	case <-time.After(3 * time.Second):
		return "result", nil
	case <-done:
		return "", fmt.Errorf("cancelled")
	}
}

// main creates a done cancellation channel and a goroutine that closes it
// after 8 seconds, then calls slowOperation. Because the operation's 3-second
// timer elapses before the 8-second cancellation, it completes successfully
// and main prints the result; otherwise it would print the cancellation error
// and return.
func main() {
	done := make(chan struct{})

	go func() {
		time.Sleep(8 * time.Second)
		close(done)
	}()

	result, err := slowOperation(done)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)
}
