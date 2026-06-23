// main.go — Teaches Go variable basics: constants, explicit var declarations,
// short variable declaration (:=), zero values for uninitialized variables,
// formatted output with fmt.Printf, and integer-to-rune conversion.
package main

import "fmt"

// greeting is a package-level string constant used as the salutation in the
// formatted output produced by main.
const greeting = "Hello"

// main demonstrates Go variable fundamentals: an explicit typed var
// declaration, short variable declaration with :=, a zero-valued float, and
// formatted printing. It also converts an int (65) to its rune (the character
// 'A') and prints it, illustrating integer-to-rune conversion.
func main() {
	var name string = "Alice"
	age := 30
	var height float64

	fmt.Printf("%s, %s. Age %d, height %.2f.\n", greeting, name, age, height)
	var i int = 65
	var c rune = rune(i)
	fmt.Printf("Rune for %d  is %c (%q)\n", i, c, c)
}
