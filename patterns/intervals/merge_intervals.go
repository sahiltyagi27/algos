// Package intervals contains interval-based patterns.
package intervals

import "sort"

// Interval represents a closed interval [Start, End].
type Interval struct {
	Start int
	End   int
}

// Merge merges overlapping intervals.
//
// Pattern:
//
//	Sort by start time, then merge into the last interval when overlapping.
//
// Time: O(n log n)
// Space: O(n)
func Merge(items []Interval) []Interval {
	if len(items) == 0 {
		return nil
	}

	sorted := append([]Interval(nil), items...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Start < sorted[j].Start
	})

	merged := []Interval{sorted[0]}
	for _, current := range sorted[1:] {
		last := &merged[len(merged)-1]
		if current.Start <= last.End {
			if current.End > last.End {
				last.End = current.End
			}
			continue
		}
		merged = append(merged, current)
	}

	return merged
}
