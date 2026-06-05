// CONCEPT 11: String Algorithms
//
// Strings are common in interviews: frequency maps, prefix/suffix logic,
// substring search, and palindrome checks.
//
// This file includes:
//   - anagram check using counts
//   - naive substring search for clarity

package algorithms

import "fmt"

func AreAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[rune]int)
	for _, ch := range a {
		counts[ch]++
	}
	for _, ch := range b {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}
	return true
}

func IndexOf(text, pattern string) int {
	if pattern == "" {
		return 0
	}
	if len(pattern) > len(text) {
		return -1
	}

	for i := 0; i <= len(text)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func RunStringDemo() {
	fmt.Printf("are anagrams listen/silent => %v\n", AreAnagrams("listen", "silent"))
	fmt.Printf("index of %q in %q => %d\n", "go", "algorithms in go", IndexOf("algorithms in go", "go"))
}
