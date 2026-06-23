// assignment.go — Day 06 Assignment: Groups words by their length into a map.
// groupByLength returns a map[int][]string where keys are word lengths and
// values are slices of words with that length. Demonstrates make() for map
// initialisation, comma-ok idiom for map lookups, and append for growing slices.
package main

import (
	"fmt"
)

// groupByLength groups the supplied words by their character length.
// It returns a map keyed by word length, where each value is a slice of the
// words having that length. The lengths' insertion is logged on first
// encounter using the comma-ok idiom to detect a missing key, and append is
// used to grow each bucket. Word order within a bucket follows input order.
func groupByLength(words []string) map[int][]string {
	returnMap := make(map[int][]string)

	for item := range words {
		if _, ok := returnMap[len(words[item])]; !ok {
			fmt.Printf("%d %q", len(words[item]), words[item])
			returnMap[len(words[item])] = make([]string, 0)
		}
		returnMap[len(words[item])] = append(returnMap[len(words[item])], words[item])
	}
	return returnMap
}

// main is the entry point. It calls groupByLength with a sample word slice and
// prints the resulting length-to-words map to standard output.
func main() {

	output := groupByLength([]string{"this", "worlds", "aaa", "bbb", "ccc"})
	fmt.Println(output)

}
