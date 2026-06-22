// Package twopointers contains two-pointer LeetCode patterns.
package twopointers

// MaxArea solves Container With Most Water.
//
// Visual:
//
//	left wall                 right wall
//	   |                         |
//	   v                         v
//	height: [1,8,6,2,5,4,8,3,7]
//
// area = min(height[left], height[right]) * width
//
// Pattern:
//
//	Start at both ends. Area is limited by the shorter wall.
//	Move the shorter pointer because moving the taller one only reduces width
//	without improving the limiting height.
//
// Time: O(n)
// Space: O(1)
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	best := 0

	for left < right {
		width := right - left
		h := minInt(height[left], height[right])
		best = maxInt(best, h*width)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return best
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
