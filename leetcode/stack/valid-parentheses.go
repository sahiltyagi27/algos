// Package stack contains stack-based LeetCode patterns.
package stack

// IsValidParentheses returns true when every opening bracket is closed in order.
//
// Pattern:
//
//	Push opening brackets on a stack.
//	For closing brackets, the top of stack must be the matching opener.
//
// Time: O(n)
// Space: O(n)
func IsValidParentheses(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var st []rune

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			st = append(st, ch)
		case ')', ']', '}':
			if len(st) == 0 || st[len(st)-1] != pairs[ch] {
				return false
			}
			st = st[:len(st)-1]
		}
	}

	return len(st) == 0
}
