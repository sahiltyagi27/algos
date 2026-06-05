# Go Algorithms Guide

> Reference for the `algos` project.
> Run all examples with `go run .`

---

## What Are Algorithms?

An algorithm is a step-by-step method for solving a problem.

Interview answer:

> An algorithm is a clear sequence of steps that takes input, processes it, and produces the expected output. We compare algorithms by correctness, time complexity, and space complexity.

---

## Big-O Quick Reference

| Big-O | Meaning | Example |
|---|---|---|
| O(1) | constant | access by index |
| O(log n) | halves the problem | binary search |
| O(n) | scans input once | linear search |
| O(n log n) | divide and merge | merge sort |
| O(n²) | nested loops | bubble sort |
| O(2ⁿ) | tries many subsets | some backtracking |

---

## 1. Searching — `algorithms/search.go`

Linear search checks every item. It works on unsorted data and costs O(n).

Binary search requires sorted data. It repeatedly halves the search space and costs O(log n).

Interview line:

> Use binary search when the input is sorted or when the answer space has a monotonic true/false condition.

---

## 2. Sorting — `algorithms/sort.go`

Sorting puts data in order so later operations become easier.

| Algorithm | Time | Space | Notes |
|---|---|---|---|
| Bubble sort | O(n²) | O(1) | learning algorithm |
| Merge sort | O(n log n) | O(n) | stable, divide-and-conquer |

In production Go, prefer:

```go
sort.Ints(nums)
sort.Slice(items, func(i, j int) bool { return items[i].Score < items[j].Score })
```

---

## 3. Recursion — `algorithms/recursion.go`

Recursion solves a problem by calling the same function on a smaller problem.

Every recursive algorithm needs:

- base case
- recursive case

Common recursion topics:

- tree traversal
- DFS
- backtracking
- divide-and-conquer sorting

---

## 4. Two Pointers — `algorithms/two_pointers.go`

Two pointers scan from two positions.

Common patterns:

- left and right ends of a sorted array
- slow and fast pointer
- read and write pointer

Examples:

- two sum in sorted array
- palindrome check
- remove duplicates

---

## 5. Sliding Window — `algorithms/sliding_window.go`

Sliding window keeps a moving range over an array or string.

Fixed-size window:

> Best sum of exactly k consecutive items.

Variable-size window:

> Longest substring without repeating characters.

Interview line:

> Sliding window usually turns an O(n²) substring/subarray problem into O(n).

---

## 6. Graph Traversal — `algorithms/graph_traversal.go`

BFS uses a queue and explores level by level.

DFS uses recursion or a stack and explores deep before backtracking.

Common graph interview questions:

- connected components
- shortest path in unweighted graph
- cycle detection
- topological sort

---

## 7. Shortest Path — `algorithms/shortest_path.go`

Dijkstra's algorithm finds shortest paths in weighted graphs with non-negative weights.

Important note:

> Dijkstra does not work with negative edge weights. For negative weights, look at Bellman-Ford.

Readable implementation in this repo:

- O(V² + E)

Production-style implementation:

- priority queue
- O((V + E) log V)

---

## 8. Dynamic Programming — `algorithms/dynamic_programming.go`

Dynamic programming stores answers to repeated subproblems.

Use DP when the problem has:

- overlapping subproblems
- optimal substructure

Classic examples:

- Fibonacci
- climb stairs
- knapsack
- longest common subsequence
- coin change

Interview approach:

1. Define the state.
2. Define the transition.
3. Define base cases.
4. Choose top-down memoization or bottom-up tabulation.

---

## 9. Greedy — `algorithms/greedy.go`

Greedy algorithms choose the best local option at each step.

They are fast but need proof that local choices produce the global best answer.

Classic examples:

- activity selection
- interval scheduling
- minimum number of coins for some coin systems
- Huffman coding

---

## 10. Backtracking — `algorithms/backtracking.go`

Backtracking tries choices, explores, then undoes the choice.

Template:

```go
choose
explore
unchoose
```

Common examples:

- permutations
- subsets
- combination sum
- Sudoku
- N-Queens

---

## 11. String Algorithms — `algorithms/string.go`

Common string interview patterns:

- frequency maps
- prefix/suffix checks
- substring search
- palindrome checks
- sliding window

This repo includes an anagram check and a clear naive substring search.

---

## Quick Interview Summary

| Algorithm Family | Use When |
|---|---|
| Linear search | data is unsorted or small |
| Binary search | data/answer space is sorted or monotonic |
| Sorting | order makes later logic easier |
| Recursion | problem naturally splits into smaller copies |
| Two pointers | sorted arrays, pairs, palindromes |
| Sliding window | contiguous subarray/substring problems |
| BFS | shortest path in unweighted graph, level order |
| DFS | traversal, components, cycle detection |
| Dijkstra | weighted shortest path with non-negative weights |
| Dynamic programming | repeated subproblems |
| Greedy | local best choice is globally correct |
| Backtracking | enumerate valid combinations |

