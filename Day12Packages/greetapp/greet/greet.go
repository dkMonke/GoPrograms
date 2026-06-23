// greet.go — Reusable greet package demonstrating exported vs unexported functions.
// Hello and Formal are exported (capitalised) and callable from other packages.
// formal (lowercase) is unexported — only accessible within the greet package.
// Formal wraps formal to expose its functionality while keeping the implementation private.
package greet

import "fmt"

// Hello returns a casual greeting addressed to name (e.g. "Hello, Alice!").
// It is exported, so it can be called from other packages.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// formal returns a formal greeting addressed to name (e.g. "Good day, Bob.").
// It is unexported (lowercase), so it is only callable within the greet package.
func formal(name string) string {
	return fmt.Sprintf("Good day, %s.", name)
}

// Formal is the exported wrapper around the unexported formal function. It lets
// callers in other packages obtain a formal greeting while keeping the
// implementation private to the greet package.
func Formal(name string) string {
	return formal(name)
}
