// main.go — Day 17: Select statement for timeout and cancellation.
// slowOperation uses select to wait on either a 3-second timer (simulating work)
// or a done channel (cancellation signal). The main goroutine closes done after
// 8 seconds. Since 3 < 8, the operation completes successfully here.
package main

import (
	"fmt"
	"time"
)

func slowOperation(done <-chan struct{}) (string, error) {
	select {
	case <-time.After(3 * time.Second):
		return "result", nil
	case <-done:
		return "", fmt.Errorf("cancelled")
	}
}

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
