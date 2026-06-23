// excersise.go — Day 03 Exercise: Practice with multi-return functions.
// divmod returns quotient, remainder, and an error for division-by-zero.
// stats is a variadic function that computes min, max, and mean of float64 values
// using named return values. Demonstrates error handling and the blank identifier (_).
package main
import "fmt"

func divmod(a,b int) (int, int, error) {
	if b == 0 {
		return 0,0,fmt.Errorf("division by zero")
	}
	return a / b, a % b, nil
	}

func stats(nums ...float64) (min, max, mean float64, err error) {
arrlength := len(nums)
if arrlength == 0 {
return 0,0,0,fmt.Errorf("No Arguments passed")
}
total :=float64(0)
min = nums[0]
max = nums[0]

for _,n := range nums {
	total +=n;
	if min < n {
		min = n
 		} 
	if max > n {
		max = n
		}
	}
mean = total /float64(arrlength)
fmt.Printf("Results are %f %f %f %s",min,max,mean,err);
return min,max,mean,err;

}

func main() {

	q, r, err := divmod(17,5)
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	fmt.Printf("17/5 = %d remainder %d\n",q,r)

	_, r2, _ := divmod(20,7)
	fmt.Printf("20 %% 7 = %d\n",r2)

	nums:=[]float64{1,2,3,4}
//	numsemp:=[]int{}
//	numsemparr:=[]string{"a","b","c","d"}
	_,_,_,_=stats(nums...)
//      _,_,_,_=stats(numsemp...)
//       _,_,_,_=stats(numsemparr...)
        
}
