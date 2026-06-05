// Package hashmap contains hash-map based LeetCode patterns.
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
