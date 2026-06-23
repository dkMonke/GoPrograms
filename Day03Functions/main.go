// main.go — Day 03: Introduction to Go functions.
// Covers basic function declarations, multiple return values (divide returns float64 + error),
// variadic functions (sum accepts any number of ints), spreading a slice with ...,
// and higher-order functions (apply takes a function as an argument).
package main

import "fmt"

// main is the program entry point and serves as a scratchpad demonstrating
// several function forms: separately typed parameters (add1), grouped
// parameter typing (add2), a multi-return function with error (divide),
// a variadic function (sum) called both directly and via slice spreading,
// and a higher-order function (apply) that accepts a function value.
func main(){

// add1 returns the sum of two int parameters declared with explicit
// individual types.
func add1(a int, b int) int {
return a+b
}

// add2 returns the sum of two int parameters using the shorthand grouped
// type declaration (a, b int).
func add2(a, b int) int {

return a+b
}

// divide returns a/b as a float64 along with an error. If b is zero it
// returns (0, error) with a "divide by zero" message; otherwise the error
// is nil.
func divide(a,b float64) (float64,error){
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
		}
	return a/b, nil

}

var c := add1(1,2)
var d := add2(2,3)
var e float64
var err string
e,err := divide(4,2)

// sum is a variadic function that returns the total of all int arguments
// passed to it. With no arguments it returns 0.
func sum(nums ...int) int {
	total := 0
	for _,n := range nums {
		total += n
	}
	return total
}
sum(1,2,3)
numbs := []int{1,2,3}
sum(nums...)

// apply is a higher-order function: it calls the supplied function f with the
// argument x and returns the result. It illustrates passing functions as
// values in Go.
func apply(f func(int) int, x int) int {
	return f(x)
}
double := func(x int) int { return x*2 }
var result := apply(double,5)
fmt.Printf("%d",result)
}
