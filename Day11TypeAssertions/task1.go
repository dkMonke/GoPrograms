// task1.go — Day 11: Type assertion fundamentals.
// Shows direct type assertion (x.(string)), the comma-ok pattern (x.(int))
// to avoid panics, and a describe function using a type switch to return
// a human-readable string for int, string, []int, and nil values.
package main

import "fmt"

func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int %d", x)
	case string:
		return fmt.Sprintf("string %q", x)
	case []int:
		return fmt.Sprintf("[]int of length %d", len(x))
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("unknown type %T", v)
	}
}

func main() {
	var x any = "hello"     //any is an empty interface which can match to any datatype present in go
	fmt.Println(x.(string)) //type assertion to string
	//fmt.Println(x.(int))  //type assertion to int which will create a panic

	//to avoid panic always use comma-ok type of check
	test, ok := x.(int)
	fmt.Println("ok", ok)
	fmt.Println("test", test)

	fmt.Println("the type of x is ", describe(x))

}
