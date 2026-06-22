// CONCEPT 1: Searching
//
// Linear search:
//   Check every item. Works on unsorted data. Time O(n).
//
// Visual:
//
//	nums:   [8, 3, 10, 1, 6]
//	         ^
//	         scan one by one until target is found
//
// Binary search:
//   Repeatedly cut a sorted search space in half. Time O(log n).
//   Requirement: data must be sorted, or the condition must be monotonic.
//
// Visual:
//
//	sorted: [1, 3, 6, 8, 10]
//	         L     M      R
//
// If target > nums[M], discard the left half:
//
//	sorted: [1, 3, 6, 8, 10]
//	                  L   R

package algorithms

import "fmt"

// LinearSearch scans left to right.
//
// Use when data is unsorted or very small.
func LinearSearch(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

// BinarySearch keeps a shrinking search window [left, right].
//
// Flow:
//
//  1. Check mid.
//  2. If target is larger, move left to mid+1.
//  3. If target is smaller, move right to mid-1.
//  4. Stop when left crosses right.
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
