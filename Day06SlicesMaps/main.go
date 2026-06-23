// main.go — Day 06: Slices and Maps.
// Builds a word-frequency map from a string slice, sorts the keys alphabetically,
// and prints counts. Also demonstrates slice aliasing — modifying a sub-slice
// (half) mutates the original slice (s) because they share underlying memory.
package main

import (
	"fmt"
	"sort"
)

// main is the entry point. It builds a word-frequency map from a sample slice,
// collects and alphabetically sorts the map keys for deterministic output, and
// prints each word with its count. It then demonstrates slice aliasing: mutating
// the first element of a sub-slice (half) also changes the backing array shared
// with the original slice (s), so both print the modified value.
func main() {

	words := []string{"the", "cat", "sat", "on", "the", "mat", "the", "cat"}
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, counts[k])
	}

	s := []int{1, 2, 3, 4, 5}
	half := s[:3]
	half[0] = 999
	fmt.Println("s:", s)
	fmt.Println("half:", half)

}
