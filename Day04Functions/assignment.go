// assignment.go — Day 04 Assignment: Counts vowels, consonants, uppercase,
// and lowercase letters in a string passed via the command line.
// Demonstrates named return values, range over a string (yields runes),
// switch for vowel detection, and the unicode package for case checks.
package main

import (
	"fmt"
	"os"
	"unicode"
)

// letterCount scans inputstring rune by rune and returns four named counts:
// the number of vowels, consonants, uppercase letters, and lowercase letters.
// Vowel detection uses a switch on the lowercase ASCII vowels, while case
// classification uses unicode.IsUpper and unicode.IsLower. As a side effect it
// prints the index and value of each rune as it iterates.
func letterCount(inputstring string) (vowels int, consonants int, ucase int, lcase int) {
	for i, v := range inputstring {
		fmt.Printf("%d location %c value", i, v)
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		default:
			consonants++
		}
		if unicode.IsUpper(v) {
			ucase++
		} else if unicode.IsLower(v) {
			lcase++
		}
	}

	return vowels, consonants, ucase, lcase
}

// main is the program entry point. It reads the first command-line argument
// as the input string, passes it to letterCount, and prints the resulting
// vowel, consonant, uppercase, and lowercase counts. It assumes at least one
// argument is provided.
func main() {
	var inputstring string
	inputstring = string(os.Args[1])
	vowels, consonants, ucase, lcase := letterCount(inputstring)
	fmt.Printf("vowels %d\n consonants %d\n ucase %d\n lcase %d\n", vowels, consonants, ucase, lcase)

}
