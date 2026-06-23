// mathlib.go — Day 13: A small math utility package designed for testing.
// Divide performs float64 division and returns a sentinel error (ErrDivByZero)
// when the divisor is zero. This clean, testable API is exercised by
// the companion mathlib_test.go file using table-driven tests.
package mathlib

import "errors"

var ErrDivByZero = errors.New("division by zero")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil

}
