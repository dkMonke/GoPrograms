// Package main demonstrates using context.Context to bound the lifetime of
// concurrent operations. It launches several simulated slow operations under a
// shared context with a timeout and shows how they are cancelled when the
// deadline is exceeded.
package main

import (
	"context"
	"fmt"
	"time"
)

// slowOp simulates a long-running operation that respects context cancellation.
//
// It takes a context ctx used to signal cancellation or timeout, and an id used
// to identify the operation in log output. The function blocks for up to two
// seconds: if that time elapses first, it prints a completion message and
// returns nil; if ctx is cancelled or its deadline is exceeded first, it returns
// an error wrapping ctx.Err() (via %w) identifying the cancelled operation.
func slowOp(ctx context.Context, id int) error {

	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("op %d done\n", id)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("op %d cancelled: %w", id, ctx.Err())

	}
}

// main creates a background context with a three-second timeout and launches
// four slowOp goroutines (ids 0 through 3) that share it. Because each slowOp
// completes in two seconds while the context deadline is three seconds, all
// operations normally finish before cancellation; the deferred cancel releases
// the context's resources. main then sleeps for five seconds to give the
// goroutines time to run and print their results before the program exits.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i <= 3; i++ {
		go func(i int) {
			err := slowOp(ctx, i)
			if err != nil {
				fmt.Println("error:", err)
			}
		}(i)
	}
	time.Sleep(5 * time.Second)
}
