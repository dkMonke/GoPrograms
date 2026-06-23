// main.go — Day 04: Go control flow constructs.
// Demonstrates if-with-initialiser (x := 42; x > 10), map literals and range iteration,
// tagless switch for grading logic, and defer to postpone execution until function return.
package main

import "fmt"

// main is the program entry point. It demonstrates several Go control-flow
// constructs: an if statement with an initialiser, building and ranging over a
// map of country-to-capital, a tagless switch implementing letter-grade logic,
// and a deferred Println that runs after the surrounding output.
func main() {
	if x := 42; x > 10 {
		fmt.Println("big", x)
	}

	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"USA":    "Washington",
	}
	for country, capital := range capitals {
		fmt.Printf("%s->%s\n", country, capital)
	}

	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C or below")
	}
	defer fmt.Println("goodbye")
	fmt.Println("hello")
}
