// assignment.go — Day 05 Assignment: Reads a key=value config file line by line.
// Demonstrates custom error types (ErrDuplicateKey), sentinel errors (InvalidKeyValueError),
// strings.Cut for splitting on "=", bufio.Scanner for line-by-line file reading,
// and map-based duplicate key detection.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ErrDuplicateKey is a custom error type returned when a configuration key
// appears more than once. It records the offending key and its value so the
// Error method can produce a descriptive message.
type ErrDuplicateKey struct {
	keyVar   string
	valueVar string
}

// InvalidKeyValueError is a sentinel error returned when a configuration line
// cannot be split into a key and a value (i.e. it contains no "=" separator).
var InvalidKeyValueError = errors.New("Invalid Key value pair in line\n")

// Error implements the error interface for *ErrDuplicateKey, returning a
// formatted message that includes the duplicate key and its value.
func (e *ErrDuplicateKey) Error() string {
	return fmt.Sprintf("Duplicate key found for key %q and value %q\n", e.keyVar, e.valueVar)
}

// readConfig parses a single configuration line of the form "key=value" and
// stores the result in the keyVal map. It returns InvalidKeyValueError when the
// line has no "=" separator, an *ErrDuplicateKey when the key already exists in
// the map, and nil on success. On success the map is updated in place as a side
// effect. It also prints the parsed components for debugging.
func readConfig(lineData string, keyVal map[string]string) error {
	before, after, found := strings.Cut(lineData, "=")
	fmt.Printf("%q %q %t", before, after, found)
	if !found {
		return InvalidKeyValueError
	}
	//val,present := keyVal[before]
	if _, present := keyVal[before]; present {
		return &ErrDuplicateKey{keyVar: before, valueVar: after}
	}
	keyVal[before] = after
	return nil
}

// main is the program entry point. If a file path is supplied as the first
// command-line argument, it opens the file, reads it line by line with a
// bufio.Scanner, and feeds each line to readConfig to build a key/value map,
// printing any errors and the final map. The file is closed via defer.
func main() {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Errorf("Error : %w", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		keyVal := make(map[string]string)
		for scanner.Scan() {
			fmt.Printf("Reading line %q\n", scanner.Text())
			err := readConfig(scanner.Text(), keyVal)
			fmt.Println(err)
			if err != nil {
				fmt.Errorf("Error %w", err)
			}

		}
		fmt.Println(keyVal)
	}

}
