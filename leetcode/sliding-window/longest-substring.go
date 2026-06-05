// Package slidingwindow contains sliding-window LeetCode patterns.
package slidingwindow

// LengthOfLongestSubstring returns the longest substring length without repeats.
//
// Pattern:
//
//	Keep a moving window [left:right].
//	If a character repeats inside the window, move left after its previous index.
//
// Time: O(n)
// Space: O(min(n, charset))
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if last, ok := lastSeen[ch]; ok && last >= left {
			left = last + 1
		}
		lastSeen[ch] = right

		if length := right - left + 1; length > best {
			best = length
		}
	}

	return best
}
