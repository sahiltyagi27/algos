// CONCEPT 3: Recursion
//
// Recursion means a function solves a problem by calling itself on a smaller
// version of the same problem.
//
// Every recursive solution needs:
//   - base case: when to stop
//   - recursive case: how to reduce the problem
//
// Visual for Factorial(4):
//
//	Factorial(4)
//	  -> 4 * Factorial(3)
//	         -> 3 * Factorial(2)
//	                -> 2 * Factorial(1)
//	                       -> 1 base case
//
// Then results return back up the call stack.

package algorithms

import "fmt"

// Factorial shows classic recursion:
//
//	base case: n <= 1
//	recursive case: n * Factorial(n-1)
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Sum recursively consumes the first element and solves the smaller slice.
//
//	Sum([2,4,6]) = 2 + Sum([4,6])
//	             = 2 + 4 + Sum([6])
//	             = 2 + 4 + 6 + Sum([])
func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + Sum(nums[1:])
}

func RunRecursionDemo() {
	fmt.Printf("factorial(5)=%d\n", Factorial(5))
	fmt.Printf("recursive sum [2 4 6]=%d\n", Sum([]int{2, 4, 6}))
}
