// Package bitmanipulation contains bit manipulation patterns.
package bitmanipulation

// SingleNumber returns the value that appears once when every other value
// appears exactly twice.
//
// Pattern:
//
//	x ^ x = 0 and x ^ 0 = x.
//
// Time: O(n)
// Space: O(1)
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func CountBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}
