// Package monotonicstack contains monotonic stack patterns.
//
// Monotonic decreasing stack visual:
//
//	temperatures: [73, 74, 75, 71]
//	stack holds indexes whose warmer day has not been found yet
//
//	When 74 appears, it pops 73 because 74 is warmer.
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
