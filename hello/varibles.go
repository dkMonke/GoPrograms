// varibles.go — Demonstrates Go variable declarations, constants, type inference
// with :=, explicit typing with var, zero values (height defaults to 0.0),
// and rune/character conversion from an integer.
package main

import "fmt"

// greeting is a package-level constant holding the greeting prefix used in
// the formatted output.
const greeting = "Hello"

// main is the program entry point. It demonstrates several variable forms:
// an explicitly typed string, type inference with :=, a float64 left at its
// zero value, and converting an int to its corresponding rune for display.
func main() {
	var name string = "Alice"
	age := 30
	var height float64
	fmt.Printf("%s, %s. Agen %d, height %.2f.\n", greeting, name, age, height)

	var i int = 65
	var c rune = rune(i)
	fmt.Printf("Rune for %d is %c (%q)\n", i, c, c)
}
