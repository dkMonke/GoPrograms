// main.go — Basic "Hello, World!" program that optionally accepts a name
// from the command-line arguments. Demonstrates os.Args for CLI input
// and fmt.Printf for formatted output.
package main

import (
	"fmt"
	"os"
)

// main is the program entry point. It greets a name taken from the first
// command-line argument, defaulting to "world" when no argument is provided.
func main() {
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s.\n", name)
}
