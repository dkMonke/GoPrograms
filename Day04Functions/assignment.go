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
func letterCount(inputstring string) (vowels int, consonants int,ucase int,lcase int) {
for i,v := range inputstring{
	fmt.Printf("%d location %c value",i,v)
	switch v {
		case 'a','e','i','o','u':
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

return vowels,consonants,ucase,lcase
}

func main() {
	var inputstring string
	inputstring = string(os.Args[1])
	vowels,consonants,ucase,lcase := letterCount(inputstring)
	fmt.Printf("vowels %d\n consonants %d\n ucase %d\n lcase %d\n",vowels,consonants,ucase,lcase)

}
