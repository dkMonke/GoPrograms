// mathlib.go — Day 13: A small math utility package designed for testing.
// Divide performs float64 division and returns a sentinel error (ErrDivByZero)
// when the divisor is zero. This clean, testable API is exercised by
// the companion mathlib_test.go file using table-driven tests.
package mathlib

import "errors"

// ErrDivByZero is the sentinel error returned by Divide when the divisor is
// zero. Callers can detect it with errors.Is(err, ErrDivByZero).
var ErrDivByZero = errors.New("division by zero")

// Divide returns the result of a divided by b. If b is zero it returns 0 and
// ErrDivByZero; otherwise it returns the quotient and a nil error.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil

}
