// CONCEPT 9: Greedy Algorithms
//
// A greedy algorithm makes the best local choice at each step.
//
// Greedy works when local choices lead to a global optimum. It is fast when
// applicable, but proving correctness matters.

package algorithms

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

// MaxNonOverlappingIntervals chooses the interval that ends earliest, then
// repeats. This is the classic activity-selection greedy strategy.
func MaxNonOverlappingIntervals(intervals []Interval) []Interval {
	sorted := append([]Interval(nil), intervals...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].End < sorted[j].End
	})

	var chosen []Interval
	lastEnd := -1 << 30
	for _, interval := range sorted {
		if interval.Start >= lastEnd {
			chosen = append(chosen, interval)
			lastEnd = interval.End
		}
	}

	return chosen
}

func RunGreedyDemo() {
	intervals := []Interval{{1, 3}, {2, 4}, {3, 5}, {0, 7}, {5, 9}}
	fmt.Printf("max non-overlapping from %v => %v\n", intervals, MaxNonOverlappingIntervals(intervals))
}
