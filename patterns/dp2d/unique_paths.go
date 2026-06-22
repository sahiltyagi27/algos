// Package dp2d contains two-dimensional dynamic programming patterns.
//
// Unique Paths visual:
//
//	Start at top-left, reach bottom-right.
//	Only moves: right or down.
//
//	+---+---+---+
//	| S |   |   |
//	+---+---+---+
//	|   |   | E |
//	+---+---+---+
//
// Ways to reach a cell:
//
//	from top + from left
package dp2d

// UniquePaths returns how many ways a robot can move from top-left to
// bottom-right in an m x n grid using only right and down moves.
//
// State:
//
//	dp[r][c] = ways to reach cell (r, c)
//
// Transition:
//
//	dp[r][c] = dp[r-1][c] + dp[r][c-1]
//
// Time: O(m*n)
// Space: O(n)
func UniquePaths(m, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[col] += dp[col-1]
		}
	}

	return dp[n-1]
}
