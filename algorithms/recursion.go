// CONCEPT 3: Recursion
//
// Recursion means a function solves a problem by calling itself on a smaller
// version of the same problem.
//
// Every recursive solution needs:
//   - base case: when to stop
//   - recursive case: how to reduce the problem

package algorithms

import "fmt"

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

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
