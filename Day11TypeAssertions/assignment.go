// assignment.go — Day 11 Assignment: Recursive pretty-printer for nested data.
// prettyPrint handles map[string]any and prettyPrintArr handles []any.
// Type switches distinguish nil, bool, int, float64, string, nested slices,
// and nested maps, printing each with appropriate formatting. Demonstrates
// recursive type assertion over heterogeneous Go data structures.
package main

import (
	"fmt"
)

// prettyPrintArr prints each element of the obj slice on its own line, using a
// type switch to format values according to their dynamic type: nil, bool, int
// (and int64), float64 (two decimals), and string. Values whose type is not
// handled are silently skipped.
func prettyPrintArr(obj []any) {
	for _, v := range obj {
		switch x := v.(type) {
		case nil:
			fmt.Println("<nil>,")
		case bool:
			if x {
				fmt.Println("true,")
			} else {
				fmt.Println("false,")
			}
		case int, int64:
			fmt.Printf("%d,\n", x)
		case float64:
			fmt.Printf("%.2f,\n", x)
		case string:
			fmt.Printf("%s,\n", x)
		}
	}
}

// prettyPrint walks the values of the obj map and prints each one formatted by
// its dynamic type via a type switch. Scalars (nil, bool, int/int64, float64,
// quoted string) print directly, while nested []any and map[string]any values
// recurse into prettyPrintArr and prettyPrint respectively. Map iteration order
// is not deterministic, so output ordering may vary between runs.
func prettyPrint(obj map[string]any) {
	for _, v := range obj {
		switch x := v.(type) {
		case nil:
			fmt.Println("<nil>,")
		case bool:
			if x {
				fmt.Println("true,")
			} else {
				fmt.Println("false,")
			}
		case int, int64:
			fmt.Printf("%d,\n", x)
		case float64:
			fmt.Printf("%.2f,\n", x)
		case string:
			fmt.Printf("%q,\n", x)
		case []any:
			fmt.Println("[")
			prettyPrintArr(x)
			fmt.Println("],")
		case map[string]any:
			fmt.Println("{")
			prettyPrint(x)
			fmt.Println("},")

		}
	}
}

// main builds a sample map containing scalars, a nested slice, and a nested
// map, then calls prettyPrint to demonstrate recursive, type-aware formatting.
func main() {

	obj := map[string]any{"a": 1, "b": []any{2, "three", nil}, "c": 3.1456, "d": map[string]any{"aa": 1, "bb": 1}}
	prettyPrint(obj)

}
