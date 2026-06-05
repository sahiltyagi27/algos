// CONCEPT 4: Two Pointers
//
// Two pointers means scanning from two positions, often from both ends.
//
// Common uses:
//   - pair sum in a sorted array
//   - palindrome checks
//   - removing duplicates

package algorithms

import "fmt"

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
