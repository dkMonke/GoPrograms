// varibles.go — Demonstrates Go variable declarations, constants, type inference
// with :=, explicit typing with var, zero values (height defaults to 0.0),
// and rune/character conversion from an integer.
package main

import "fmt"

const greeting = "Hello"

func main() {
	var name string = "Alice"
	age := 30
	var height float64
	fmt.Printf("%s, %s. Agen %d, height %.2f.\n",greeting,name,age,height)

	var i int = 65
	var c rune = rune(i)
	fmt.Printf("Rune for %d is %c (%q)\n",i,c,c)
}
