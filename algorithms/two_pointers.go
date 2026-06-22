// CONCEPT 4: Two Pointers
//
// Two pointers means scanning from two positions, often from both ends.
//
// Visual:
//
//	nums: [1, 2, 4, 7, 11]
//	       L              R
//
// If sum is too small, move L right.
// If sum is too large, move R left.
//
// Common uses:
//   - pair sum in a sorted array
//   - palindrome checks
//   - removing duplicates

package algorithms

import "fmt"

// TwoSumSorted works because the array is sorted.
//
// When sum is smaller than target, increasing left can increase the sum.
// When sum is larger than target, decreasing right can decrease the sum.
func TwoSumSorted(nums []int, target int) (int, int, bool) {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		switch {
		case sum == target:
			return left, right, true
		case sum < target:
			left++
		default:
			right--
		}
	}

	return -1, -1, false
}

// IsPalindrome compares characters from both ends.
//
//	level
//	^   ^
//	L   R
//
// Move inward while characters match.
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func RunTwoPointersDemo() {
	nums := []int{1, 2, 4, 7, 11}
	i, j, ok := TwoSumSorted(nums, 9)
	fmt.Printf("two sum sorted %v target=9 => indexes=(%d,%d) found=%v\n", nums, i, j, ok)
	fmt.Printf("is palindrome %q => %v\n", "level", IsPalindrome("level"))
}
