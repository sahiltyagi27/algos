// CONCEPT 5: Sliding Window
//
// Sliding window keeps a moving range over an array/string.
//
// Visual:
//
//	nums: [2, 1, 5, 1, 3, 2], k=3
//	       [2, 1, 5]
//	          [1, 5, 1]
//	             [5, 1, 3]
//	                [1, 3, 2]
//
// Fixed-size window:
//   Useful for "best sum of k consecutive elements".
//
// Variable-size window:
//   Useful for "longest substring without repeating characters".

package algorithms

import "fmt"

// MaxSumSubarray keeps the sum of exactly k items.
//
// Slide by:
//
//	add new right value
//	remove value that left the window
func MaxSumSubarray(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	best := windowSum
	for right := k; right < len(nums); right++ {
		windowSum += nums[right]
		windowSum -= nums[right-k]
		if windowSum > best {
			best = windowSum
		}
	}

	return best
}

// LongestUniqueSubstring uses a variable-size window.
//
// Window:
//
//	left ... right
//
// If s[right] was last seen inside the current window, move left after the
// previous occurrence.
func LongestUniqueSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if last, ok := lastSeen[ch]; ok && last >= left {
			left = last + 1
		}
		lastSeen[ch] = right
		if length := right - left + 1; length > best {
			best = length
		}
	}

	return best
}

func RunSlidingWindowDemo() {
	nums := []int{2, 1, 5, 1, 3, 2}
	fmt.Printf("max sum subarray k=3 in %v => %d\n", nums, MaxSumSubarray(nums, 3))
	fmt.Printf("longest unique substring in %q => %d\n", "abcabcbb", LongestUniqueSubstring("abcabcbb"))
}
