// algos: Small Go algorithm examples for interview preparation.
//
// Run:
//   go run .

package main

import (
	"fmt"

	"algos/algorithms"
)

func main() {
	fmt.Println("=== Searching ===")
	algorithms.RunSearchDemo()

	fmt.Println("\n=== Sorting ===")
	algorithms.RunSortDemo()

	fmt.Println("\n=== Recursion ===")
	algorithms.RunRecursionDemo()

	fmt.Println("\n=== Two Pointers ===")
	algorithms.RunTwoPointersDemo()

	fmt.Println("\n=== Sliding Window ===")
	algorithms.RunSlidingWindowDemo()

	fmt.Println("\n=== Graph Traversal ===")
	algorithms.RunGraphTraversalDemo()

	fmt.Println("\n=== Shortest Path ===")
	algorithms.RunShortestPathDemo()

	fmt.Println("\n=== Dynamic Programming ===")
	algorithms.RunDynamicProgrammingDemo()

	fmt.Println("\n=== Greedy ===")
	algorithms.RunGreedyDemo()

	fmt.Println("\n=== Backtracking ===")
	algorithms.RunBacktrackingDemo()

	fmt.Println("\n=== String Algorithms ===")
	algorithms.RunStringDemo()
}
