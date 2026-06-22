# Go Algorithms

Small Go algorithm examples for interview preparation.

## DSA Interview Coding Rule

For live coding, keep Go code vanilla and writable from memory.

```text
Queue   = []int, append to push, queue[0] + queue = queue[1:] to pop
Stack   = []int, append to push, stack[len(stack)-1] + stack = stack[:len(stack)-1] to pop
Set     = map[T]bool, or []bool for 0..n-1 nodes
Graph   = [][]int for 0..n-1 nodes
Visited = []bool for 0..n-1 nodes
```

Avoid custom `Stack`, `Queue`, or `Set` structs in DSA interview solutions unless the problem explicitly asks you to implement that data structure.

## Run

```bash
go run .
go test ./...
```

## Topic Index

| Topic | File |
|---|---|
| Core algorithms folder | [algorithms/README.md](algorithms/README.md) |
| Linear search and binary search | [algorithms/search.go](algorithms/search.go) |
| Bubble sort and merge sort | [algorithms/sort.go](algorithms/sort.go) |
| Recursion | [algorithms/recursion.go](algorithms/recursion.go) |
| Two pointers | [algorithms/two_pointers.go](algorithms/two_pointers.go) |
| Sliding window | [algorithms/sliding_window.go](algorithms/sliding_window.go) |
| BFS and DFS | [algorithms/graph_traversal.go](algorithms/graph_traversal.go) |
| Grid DFS and surrounded city components | [algorithms/graph_traversal.go](algorithms/graph_traversal.go) |
| Dijkstra shortest path | [algorithms/shortest_path.go](algorithms/shortest_path.go) |
| Topological sort / Course Schedule | [patterns/topologicalsort/README.md](patterns/topologicalsort/README.md) |
| Dynamic programming | [algorithms/dynamic_programming.go](algorithms/dynamic_programming.go) |
| Greedy algorithms | [algorithms/greedy.go](algorithms/greedy.go) |
| Backtracking | [algorithms/backtracking.go](algorithms/backtracking.go) |
| String algorithms | [algorithms/string.go](algorithms/string.go) |

## Main Guide

Read the full study notes here:

- [GUIDE.md](GUIDE.md)
- [NEETCODE_ROADMAP.md](NEETCODE_ROADMAP.md)

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

## Pattern Coverage

Additional NeetCode-style patterns live in:

- [patterns/README.md](patterns/README.md)
- [patterns/prefixsum](patterns/prefixsum)
- [patterns/intervals](patterns/intervals)
- [patterns/trie](patterns/trie)
- [patterns/monotonicstack](patterns/monotonicstack)
- [patterns/unionfind](patterns/unionfind)
- [patterns/topologicalsort](patterns/topologicalsort/README.md)
- [patterns/dp2d](patterns/dp2d)
- [patterns/bitmanipulation](patterns/bitmanipulation)
- [patterns/mathgeometry](patterns/mathgeometry)
- [patterns/binarysearchanswer](patterns/binarysearchanswer)
- [patterns/advancedgraphs](patterns/advancedgraphs)

## Interview Highlights

- Binary search needs sorted data or a monotonic condition.
- Sliding window usually optimizes contiguous subarray or substring problems.
- BFS is natural for shortest paths in unweighted graphs.
- DFS is natural for components, recursion, and backtracking.
- For grid DFS, clarify whether the interviewer means individual cells or connected components.
- Dynamic programming stores answers to repeated subproblems.
- Greedy works only when local best choices lead to the global best answer.
