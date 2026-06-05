// CONCEPT 8: Dynamic Programming
//
// Dynamic programming solves problems with:
//   - overlapping subproblems
//   - optimal substructure
//
// Instead of recomputing the same answer many times, store previous results.

package algorithms

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2
	for step := 3; step <= n; step++ {
		current := prev1 + prev2
		prev2, prev1 = prev1, current
	}
	return prev1
}

func RunDynamicProgrammingDemo() {
	fmt.Printf("fibonacci(8)=%d\n", Fibonacci(8))
	fmt.Printf("climb stairs n=5 => %d ways\n", ClimbStairs(5))
}
