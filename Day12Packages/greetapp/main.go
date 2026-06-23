// main.go — Entry point for greetapp. Imports the local greet package
// and calls its exported functions Hello and Formal to demonstrate
// how Go modules organise code into reusable packages.
package main

import (
	"fmt"
	"greetapp/greet"
)

func main() {
	fmt.Println(greet.Hello("Alice"))
	fmt.Println(greet.Formal("Bob"))
}
