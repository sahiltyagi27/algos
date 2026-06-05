// CONCEPT 2: Sorting
//
// Sorting arranges data in a useful order.
//
// Bubble sort:
//   Simple learning algorithm. Time O(n^2). Rarely used in production.
//
// Merge sort:
//   Divide-and-conquer algorithm. Time O(n log n). Uses extra memory O(n).

package algorithms

import "fmt"

func BubbleSort(nums []int) []int {
	out := append([]int(nil), nums...)

	for i := 0; i < len(out); i++ {
		swapped := false
		for j := 0; j < len(out)-1-i; j++ {
			if out[j] > out[j+1] {
				out[j], out[j+1] = out[j+1], out[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return out
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return append([]int(nil), nums...)
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func RunSortDemo() {
	nums := []int{5, 1, 9, 3, 7}
	fmt.Printf("bubble sort %v => %v\n", nums, BubbleSort(nums))
	fmt.Printf("merge sort  %v => %v\n", nums, MergeSort(nums))
}
