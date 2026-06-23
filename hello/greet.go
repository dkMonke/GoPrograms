// greet.go — Greets each name passed as a command-line argument.
// If no arguments are given, prints "Hello, world." as the default greeting.
// Demonstrates looping over os.Args and formatted printing with fmt.Printf.
package main

import (
	"fmt"
	"os"
	)

func main() {
	fmt.Printf("%d",len(os.Args))
	if len(os.Args) <= 1 {
	fmt.Println("Hello, world.")
	}
	for i := 1; i<len(os.Args); i++ {
	fmt.Printf("Hello, %s.\n",os.Args[i])
	}
}


