// mathlib_test.go — Day 13: Table-driven unit tests for Divide.
// Uses Go's testing package with t.Run sub-tests. Each test case specifies
// inputs, expected result, and expected error. errors.Is checks sentinel errors,
// and t.Fatalf / t.Errorf report failures. Run with: go test ./...
package mathlib

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr error
	}{
		{"simple", 10, 2, 5, nil},
		{"fractional", 1, 3, 1.0 / 3, nil},
		{"by zero", 10, 0, 0, ErrDivByZero},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Divide(c.a, c.b)
			if !errors.Is(err, c.wantErr) {
				t.Fatalf("err: got %v, want %v", err, c.wantErr)
			}
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}
