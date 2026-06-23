// excersise.go — Day 03 Exercise: Practice with multi-return functions.
// divmod returns quotient, remainder, and an error for division-by-zero.
// stats is a variadic function that computes min, max, and mean of float64 values
// using named return values. Demonstrates error handling and the blank identifier (_).
package main

import "fmt"

// divmod divides a by b and returns the integer quotient, the remainder,
// and an error. If b is zero it returns (0, 0, error) with a
// "division by zero" message rather than panicking. On success the error
// is nil and the results are a/b and a%b.
func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	return a / b, a % b, nil
}

// stats is a variadic function that computes the minimum, maximum, and mean
// of the supplied float64 values using named return values. If no arguments
// are passed it returns zero values and a non-nil error. As a side effect it
// prints the computed results to standard output before returning them.
func stats(nums ...float64) (min, max, mean float64, err error) {
	arrlength := len(nums)
	if arrlength == 0 {
		return 0, 0, 0, fmt.Errorf("No Arguments passed")
	}
	total := float64(0)
	min = nums[0]
	max = nums[0]

	for _, n := range nums {
		total += n
		if min < n {
			min = n
		}
		if max > n {
			max = n
		}
	}
	mean = total / float64(arrlength)
	fmt.Printf("Results are %f %f %f %s", min, max, mean, err)
	return min, max, mean, err

}

// main is the program entry point. It exercises divmod with sample inputs
// (including a use of the blank identifier to discard unwanted results) and
// invokes stats by spreading a float64 slice with the ... operator.
func main() {

	q, r, err := divmod(17, 5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("17/5 = %d remainder %d\n", q, r)

	_, r2, _ := divmod(20, 7)
	fmt.Printf("20 %% 7 = %d\n", r2)

	nums := []float64{1, 2, 3, 4}
	//	numsemp:=[]int{}
	//	numsemparr:=[]string{"a","b","c","d"}
	_, _, _, _ = stats(nums...)
	//      _,_,_,_=stats(numsemp...)
	//       _,_,_,_=stats(numsemparr...)

}
