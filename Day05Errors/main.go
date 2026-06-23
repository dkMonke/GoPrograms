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

// ErrInvalidInput is a sentinel error indicating that the input string was
// empty or otherwise invalid. Callers can detect it with errors.Is.
var ErrInvalidInput = errors.New("Invalid input")

// ParseError is a custom error type describing a parsing failure. It records
// the input being parsed and the position (Pos) within it where the error
// occurred so callers can extract details via errors.As.
type ParseError struct {
	Input string
	Pos   int
}

// Error implements the error interface for *ParseError, returning a message
// that includes the failing position and the input string.
func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at %d in %q", e.Pos, e.Input)
}

// parse interprets the string s and returns its length on success. It returns
// ErrInvalidInput when s is empty and a *ParseError when s equals "bad",
// demonstrating both sentinel and custom-typed error returns.
func parse(s string) (int, error) {
	if s == "" {
		return 0, ErrInvalidInput
	}
	if s == "bad" {
		return 0, &ParseError{Input: s, Pos: 0}
	}
	return len(s), nil
}

// process calls parse on s and, if parsing fails, wraps the underlying error
// with additional context using fmt.Errorf and the %w verb so the original
// error remains retrievable via errors.Is/errors.As. It returns nil on success.
func process(s string) error {
	_, err := parse(s)
	if err != nil {
		return fmt.Errorf("processing %q: %w", s, err)
	}
	return nil
}

// main is the program entry point. It runs process over a set of sample inputs
// and classifies any returned errors using errors.Is (for the ErrInvalidInput
// sentinel) and errors.As (for the *ParseError type). It also opens a
// nonexistent file to demonstrate matching the os.ErrNotExist sentinel.
func main() {
	inputs := []string{"hello", "", "bad"}
	for _, in := range inputs {
		err := process(in)
		if err == nil {
			fmt.Println("ok:", in)
			continue
		}

		//sentinel check
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("got invalid input:", err)
			continue
		}

		// Type extraction
		var pe *ParseError
		if errors.As(err, &pe) {
			fmt.Printf("parse problem at pos %d:%v\n", pe.Pos, err)
			continue
		}
		fmt.Println("unknown error:", err)
	}
	_, err := os.Open("/nonexistent")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file not found, as expected")
	}
}
