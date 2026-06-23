// assignment.go — Day 05 Assignment: Reads a key=value config file line by line.
// Demonstrates custom error types (ErrDuplicateKey), sentinel errors (InvalidKeyValueError),
// strings.Cut for splitting on "=", bufio.Scanner for line-by-line file reading,
// and map-based duplicate key detection.
package main

import (
	"fmt"
	"os"
	"errors"
	"bufio"
	"strings"
	)

type ErrDuplicateKey struct {
	keyVar string
	valueVar string

}
var InvalidKeyValueError = errors.New("Invalid Key value pair in line\n")
func (e *ErrDuplicateKey) Error() string {
	return fmt.Sprintf("Duplicate key found for key %q and value %q\n",e.keyVar,e.valueVar)
}

func readConfig(lineData string, keyVal map[string]string) error {
	before,after,found := strings.Cut(lineData,"=")
	fmt.Printf("%q %q %t",before,after,found)
	if !found {
	return InvalidKeyValueError
	}
	//val,present := keyVal[before]
	if _,present :=  keyVal[before];present {
		return &ErrDuplicateKey{keyVar:before,valueVar:after}
 	}
	keyVal[before] = after
	return nil
}
func main() {
	if len(os.Args) > 1 {
	file,err := os.Open(os.Args[1])
	if err != nil {
	fmt.Errorf("Error : %w",err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	keyVal := make(map[string]string)	
	for scanner.Scan() {
		fmt.Printf("Reading line %q\n",scanner.Text())
		err := readConfig(scanner.Text(),keyVal)
		fmt.Println(err)
		if err != nil {
			fmt.Errorf("Error %w",err)
		}
	
	}
	fmt.Println(keyVal)
        }


}
