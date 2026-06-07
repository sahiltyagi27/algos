// Package mathgeometry contains math and geometry patterns.
package mathgeometry

// GCD returns the greatest common divisor using Euclid's algorithm.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

// IsRectangleOverlap returns true if two axis-aligned rectangles overlap.
//
// Rect format:
//
//	[x1, y1, x2, y2]
//
// Time: O(1)
// Space: O(1)
func IsRectangleOverlap(a, b [4]int) bool {
	return a[0] < b[2] &&
		a[2] > b[0] &&
		a[1] < b[3] &&
		a[3] > b[1]
}
