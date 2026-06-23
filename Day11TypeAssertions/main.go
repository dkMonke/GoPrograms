// main.go — Day 11: Type Assertions and Type Switches.
// sum is a variadic function accepting any types. A type switch (v.(type))
// inspects each argument at runtime to convert int and float64 to a running total,
// returning an error for unsupported types like string.
package main

import (
	"fmt"
)

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

func main() {
	n, err := sum(1, 2.5, 3, 4.5)
	fmt.Println(n, err)

	_, err = sum(1, "two", 3)
	fmt.Println(err)
}
