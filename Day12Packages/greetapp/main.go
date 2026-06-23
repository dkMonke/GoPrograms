// main.go — Entry point for greetapp. Imports the local greet package
// and calls its exported functions Hello and Formal to demonstrate
// how Go modules organise code into reusable packages.
package main

import (
	"fmt"
	"greetapp/greet"
)

// main is the program entry point. It calls the greet package's exported
// Hello and Formal functions and prints their casual and formal greetings,
// demonstrating cross-package function calls within a Go module.
func main() {
	fmt.Println(greet.Hello("Alice"))
	fmt.Println(greet.Formal("Bob"))
}
