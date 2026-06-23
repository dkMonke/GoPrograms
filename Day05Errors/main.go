// main.go — Day 05: Error handling patterns in Go.
// Covers sentinel errors (ErrInvalidInput), custom error types (ParseError with Error() method),
// error wrapping with fmt.Errorf and %w, unwrapping with errors.Is and errors.As,
// and checking os-level errors like os.ErrNotExist.
package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrInvalidInput = errors.New("Invalid input")

type ParseError struct {
	Input string
	Pos int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at %d in %q", e.Pos, e.Input)
}

func parse(s string) (int,error) {
	if s == "" {
		return 0,ErrInvalidInput
	}
	if s == "bad" {
		return 0, &ParseError{Input:s, Pos:0}
	}
	return len(s), nil
}

func process(s string) error {
	_, err := parse(s)
	if err != nil {
		return fmt.Errorf("processing %q: %w",s,err)
	}
	return nil
}

func main() {
	inputs := []string{"hello","","bad"}
	for _,in := range inputs {
		err := process(in)
		if err == nil {
			fmt.Println("ok:",in)
			continue
		}

		//sentinel check
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("got invalid input:",err)
			continue
		}

		// Type extraction
		var pe *ParseError
		if errors.As(err, &pe) {
			fmt.Printf("parse problem at pos %d:%v\n", pe.Pos,err)
			continue
		}
		fmt.Println("unknown error:",err)
	}
	_,err := os.Open("/nonexistent")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file not found, as expected")
	}
}
