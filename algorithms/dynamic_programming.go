// CONCEPT 8: Dynamic Programming
//
// Dynamic programming solves problems with:
//   - overlapping subproblems
//   - optimal substructure
//
// Instead of recomputing the same answer many times, store previous results.
//
// Fibonacci visual:
//
//	F(5)
//	= F(4) + F(3)
//
// Without DP, F(3), F(2), etc. are recomputed many times.
//
// DP table:
//
//	index: 0  1  2  3  4  5
//	dp:    0  1  1  2  3  5

package algorithms

import "fmt"

// Fibonacci uses bottom-up tabulation.
//
// dp[i] means Fibonacci number at i.
// Transition:
//
//	dp[i] = dp[i-1] + dp[i-2]
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

// ClimbStairs is Fibonacci-like.
//
// ways[n] = ways[n-1] + ways[n-2]
//
// To reach step n, the last move came from:
//
//	n-1 using 1 step
//	n-2 using 2 steps
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
