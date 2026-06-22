// Package prefixsum contains prefix-sum patterns.
//
// Prefix sum visual:
//
//	nums:   [2, 4, 6]
//	prefix: [0, 2, 6, 12]
//
// Sum from index 1 to 2:
//
//	prefix[3] - prefix[1] = 12 - 2 = 10
package prefixsum

// RangeSum builds prefix sums and returns the sum from left to right inclusive.
//
// Pattern:
//
//	prefix[i] stores sum of nums[0:i].
//	sum(left, right) = prefix[right+1] - prefix[left].
//
// Time:
//
//	Build O(n), query O(1)
//
// Space:
//
//	O(n)
func RangeSum(nums []int, left, right int) int {
	prefix := make([]int, len(nums)+1)
	for i, n := range nums {
		prefix[i+1] = prefix[i] + n
	}
	return prefix[right+1] - prefix[left]
}

// SubarraySumEqualsK counts subarrays whose sum is k.
//
// Pattern:
//
//	If currentPrefix - oldPrefix == k, the range between them sums to k.
//
// Time: O(n)
// Space: O(n)
func SubarraySumEqualsK(nums []int, k int) int {
	counts := map[int]int{0: 1}
	prefix := 0
	total := 0

	for _, n := range nums {
		prefix += n
		total += counts[prefix-k]
		counts[prefix]++
	}

	return total
}
