// main.go — Day 11: Type Assertions and Type Switches.
// sum is a variadic function accepting any types. A type switch (v.(type))
// inspects each argument at runtime to convert int and float64 to a running total,
// returning an error for unsupported types like string.
package main

import (
	"fmt"
)

// sum adds together a variadic list of values of any type, returning the
// running total as a float64. A type switch handles int and float64 operands;
// encountering a string returns a descriptive error, and any other unsupported
// type returns an error naming its index and dynamic type. On any error the
// returned total is 0.
func sum(value ...any) (float64, error) {
	var total float64
	for i, v := range value {
		switch x := v.(type) {
		case int:
			total += float64(x)
		case float64:
			total += x
		case string:
			return 0, fmt.Errorf("value %d is a string, not numeric", v)
		default:
			return 0, fmt.Errorf("value %d has unsupported type %T", i, v)
		}
	}
	return total, nil
}

// main exercises sum twice: once with only numeric values (succeeding) and
// once including a string (failing), printing the results and errors to show
// both the happy path and error handling.
func main() {
	n, err := sum(1, 2.5, 3, 4.5)
	fmt.Println(n, err)

	_, err = sum(1, "two", 3)
	fmt.Println(err)
}
