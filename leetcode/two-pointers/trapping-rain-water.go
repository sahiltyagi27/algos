package twopointers

// Trap solves Trapping Rain Water.
//
// Visual:
//
//	height: [0,1,0,2]
//	          _
//	     _   | |
//	 _  | | water can sit above lower bars
//
// Pattern:
//
//	Water at a position depends on the smaller of max wall on the left and
//	max wall on the right. The two-pointer version keeps leftMax and rightMax
//	while moving the side with the smaller current boundary.
//
// Time: O(n)
// Space: O(1)
func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}
