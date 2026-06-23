// greet.go — Reusable greet package demonstrating exported vs unexported functions.
// Hello and Formal are exported (capitalised) and callable from other packages.
// formal (lowercase) is unexported — only accessible within the greet package.
// Formal wraps formal to expose its functionality while keeping the implementation private.
package greet

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func formal(name string) string {
	return fmt.Sprintf("Good day, %s.", name)
}

func Formal(name string) string {
	return formal(name)
}
