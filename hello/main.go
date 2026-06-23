// main.go — Basic "Hello, World!" program that optionally accepts a name
// from the command-line arguments. Demonstrates os.Args for CLI input
// and fmt.Printf for formatted output.
package main

import (
	"fmt"
	"os"
)

func main() {
	name := "world"
	if len(os.Args) >1 {
	name = os.Args[1]
	}
	fmt.Printf("Hello, %s.\n",name)
}

