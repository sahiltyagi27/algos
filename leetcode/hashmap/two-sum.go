// Package hashmap contains hash-map based LeetCode patterns.
//
// Two Sum visual:
//
//	nums:   [2, 7, 11, 15]
//	target: 9
//
//	At 2:
//	  need 7, map is empty -> store 2 -> index 0
//
//	At 7:
//	  need 2, map has 2 -> answer [0, 1]
package hashmap

// TwoSum returns indexes of two numbers that add up to target.
//
// Pattern:
//
//	Store previously seen numbers in a map from value -> index.
//	For each number, check whether target-number was already seen.
//
// Time: O(n)
// Space: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}

	return nil
}
