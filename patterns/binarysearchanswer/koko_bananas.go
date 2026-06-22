// Package binarysearchanswer contains binary-search-on-answer patterns.
//
// Binary search on answer visual:
//
//	possible speeds: 1 ... maxPile
//
//	If speed=mid works, try smaller speeds.
//	If speed=mid fails, try larger speeds.
//
// This works because the condition is monotonic:
//
//	if speed x works, any speed greater than x also works.
package binarysearchanswer

// MinEatingSpeed solves the Koko Eating Bananas pattern.
//
// Pattern:
//
//	Search the smallest speed that satisfies canFinish(speed).
//	If speed works, try smaller. If not, try larger.
//
// Time: O(n log max(piles))
// Space: O(1)
func MinEatingSpeed(piles []int, h int) int {
	left, right := 1, maxPile(piles)

	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canFinish(piles []int, h, speed int) bool {
	hours := 0
	for _, pile := range piles {
		hours += (pile + speed - 1) / speed
	}
	return hours <= h
}

func maxPile(piles []int) int {
	maximum := 0
	for _, pile := range piles {
		if pile > maximum {
			maximum = pile
		}
	}
	return maximum
}
