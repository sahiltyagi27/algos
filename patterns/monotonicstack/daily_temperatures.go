// Package monotonicstack contains monotonic stack patterns.
package monotonicstack

// DailyTemperatures returns how many days until a warmer temperature.
//
// Pattern:
//
//	Keep indexes in a decreasing stack. When a warmer temperature appears,
//	pop colder indexes and compute the distance.
//
// Time: O(n)
// Space: O(n)
func DailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	var stack []int

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prev] = i - prev
		}
		stack = append(stack, i)
	}

	return answer
}
