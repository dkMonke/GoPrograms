// main.go — Day 03: Introduction to Go functions.
// Covers basic function declarations, multiple return values (divide returns float64 + error),
// variadic functions (sum accepts any number of ints), spreading a slice with ...,
// and higher-order functions (apply takes a function as an argument).
package main

import "fmt"

func main(){

func add1(a int, b int) int {
return a+b
}

func add2(a, b int) int {

return a+b
}

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

func apply(f func(int) int, x int) int {
	return f(x)
}
double := func(x int) int { return x*2 }
var result := apply(double,5)
fmt.Printf("%d",result)
}
