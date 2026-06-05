# Go Algorithms

Small Go algorithm examples for interview preparation.

## Run

```bash
go run .
go test ./...
```

## Topic Index

| Topic | File |
|---|---|
| Linear search and binary search | [algorithms/search.go](algorithms/search.go) |
| Bubble sort and merge sort | [algorithms/sort.go](algorithms/sort.go) |
| Recursion | [algorithms/recursion.go](algorithms/recursion.go) |
| Two pointers | [algorithms/two_pointers.go](algorithms/two_pointers.go) |
| Sliding window | [algorithms/sliding_window.go](algorithms/sliding_window.go) |
| BFS and DFS | [algorithms/graph_traversal.go](algorithms/graph_traversal.go) |
| Dijkstra shortest path | [algorithms/shortest_path.go](algorithms/shortest_path.go) |
| Dynamic programming | [algorithms/dynamic_programming.go](algorithms/dynamic_programming.go) |
| Greedy algorithms | [algorithms/greedy.go](algorithms/greedy.go) |
| Backtracking | [algorithms/backtracking.go](algorithms/backtracking.go) |
| String algorithms | [algorithms/string.go](algorithms/string.go) |

## Main Guide

Read the full study notes here:

- [GUIDE.md](GUIDE.md)

## LeetCode Practice

Pattern-based solved examples live in:

- [leetcode/README.md](leetcode/README.md)

Current examples:

| Pattern | Problem | File |
|---|---|---|
| Hash map | Two Sum | [leetcode/hashmap/two-sum.go](leetcode/hashmap/two-sum.go) |
| Stack | Valid Parentheses | [leetcode/stack/valid-parentheses.go](leetcode/stack/valid-parentheses.go) |
| Sliding window | Longest Substring Without Repeating Characters | [leetcode/sliding-window/longest-substring.go](leetcode/sliding-window/longest-substring.go) |
| Two pointers | Container With Most Water | [leetcode/two-pointers/container-with-most-water.go](leetcode/two-pointers/container-with-most-water.go) |
| Two pointers | Trapping Rain Water | [leetcode/two-pointers/trapping-rain-water.go](leetcode/two-pointers/trapping-rain-water.go) |

## Interview Highlights

- Binary search needs sorted data or a monotonic condition.
- Sliding window usually optimizes contiguous subarray or substring problems.
- BFS is natural for shortest paths in unweighted graphs.
- DFS is natural for components, recursion, and backtracking.
- Dynamic programming stores answers to repeated subproblems.
- Greedy works only when local best choices lead to the global best answer.
