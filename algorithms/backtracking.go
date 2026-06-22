// CONCEPT 10: Backtracking
//
// Backtracking explores choices one by one. If a choice cannot lead to a valid
// answer, undo it and try the next choice.
//
// Common uses:
//   - permutations
//   - subsets
//   - combination sum
//   - Sudoku / N-Queens
//
// Template:
//
//	choose
//	explore
//	unchoose
//
// Permutations visual for [1,2]:
//
//	        []
//	      /    \
//	    [1]    [2]
//	    /        \
//	 [1,2]     [2,1]

package algorithms

import "fmt"

// Permutations tries every unused number at each position.
//
// path is the current partial answer.
// used prevents reusing the same index twice.
func Permutations(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	var path []int

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			result = append(result, append([]int(nil), path...))
			return
		}

		for i, n := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, n)

			dfs()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()
	return result
}

func RunBacktrackingDemo() {
	nums := []int{1, 2, 3}
	fmt.Printf("permutations of %v => %v\n", nums, Permutations(nums))
}
