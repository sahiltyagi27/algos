// CONCEPT 1: Searching
//
// Linear search:
//   Check every item. Works on unsorted data. Time O(n).
//
// Binary search:
//   Repeatedly cut a sorted search space in half. Time O(log n).
//   Requirement: data must be sorted, or the condition must be monotonic.

package algorithms

import "fmt"

func LinearSearch(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}

	return -1
}

func RunSearchDemo() {
	unsorted := []int{8, 3, 10, 1, 6}
	sorted := []int{1, 3, 6, 8, 10}
	fmt.Printf("linear search 10 in %v => index %d\n", unsorted, LinearSearch(unsorted, 10))
	fmt.Printf("binary search 8 in %v => index %d\n", sorted, BinarySearch(sorted, 8))
}
