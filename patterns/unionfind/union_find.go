// Package unionfind contains disjoint-set union patterns.
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

func New(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

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
