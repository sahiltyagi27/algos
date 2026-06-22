// Package bitmanipulation contains bit manipulation patterns.
//
// Bit manipulation visual:
//
//	XOR:
//	  x ^ x = 0
//	  x ^ 0 = x
//
// Power of two:
//
//	8      = 1000
//	8 - 1  = 0111
//	8 & 7  = 0000
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

// IsPowerOfTwo works because powers of two have exactly one set bit.
//
// n & (n-1) clears the lowest set bit.
// For powers of two, clearing the only set bit gives 0.
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// CountBits uses the relation:
//
//	bits[i] = bits[i/2] + lastBit
//
// i>>1 removes the last bit.
// i&1 reads the last bit.
func CountBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}
