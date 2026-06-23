// main.go — Teaches Go variable basics: constants, explicit var declarations,
// short variable declaration (:=), zero values for uninitialized variables,
// formatted output with fmt.Printf, and integer-to-rune conversion.
package main
import "fmt"

const greeting = "Hello"

func main() {
	var name string = "Alice"
	age := 30
	var height float64

	fmt.Printf("%s, %s. Age %d, height %.2f.\n",greeting,name,age,height)
	var i int = 65
	var c rune = rune(i)
	fmt.Printf("Rune for %d  is %c (%q)\n",i,c,c)
}
