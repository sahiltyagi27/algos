// Package unionfind contains disjoint-set union patterns.
//
// Union-Find visual:
//
//	Before union(1,2):
//	1    2    3
//
//	After union(1,2):
//	1 -- 2    3
//
// Find(x) returns the representative/root of x's group.
package unionfind

// UnionFind tracks connected components.
//
// Common uses:
//   - connected components
//   - cycle detection in undirected graphs
//   - Kruskal minimum spanning tree
type UnionFind struct {
	parent []int
	rank   []int
}

// New starts every node as its own parent.
//
//	parent[i] = i
func New(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find returns the root representative.
//
// Path compression makes future finds faster:
//
//	1 -> 2 -> 3
//
// becomes:
//
//	1 -> 3
//	2 -> 3
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union connects two components.
//
// It returns false when a and b were already connected.
func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)
	if rootA == rootB {
		return false
	}

	if uf.rank[rootA] < uf.rank[rootB] {
		uf.parent[rootA] = rootB
	} else if uf.rank[rootA] > uf.rank[rootB] {
		uf.parent[rootB] = rootA
	} else {
		uf.parent[rootB] = rootA
		uf.rank[rootA]++
	}
	return true
}

func (uf *UnionFind) Connected(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}
