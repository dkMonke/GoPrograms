// task1.go — Day 11: Type assertion fundamentals.
// Shows direct type assertion (x.(string)), the comma-ok pattern (x.(int))
// to avoid panics, and a describe function using a type switch to return
// a human-readable string for int, string, []int, and nil values.
package main

import "fmt"

// describe returns a human-readable description of the dynamic type and value
// held in v. It uses a type switch to recognize int, string, []int (reporting
// its length), and nil, falling back to a message that names the unknown type
// via %T for anything else.
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

// main demonstrates type assertions on an empty-interface value: a direct
// assertion to string, the comma-ok form to safely test an int assertion
// without panicking, and a call to describe to report the value's type.
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
