// Package advancedgraphs contains advanced graph patterns.
//
// Bellman-Ford visual:
//
//	Relax every edge V-1 times.
//
//	Edge: u --w--> v
//
//	if dist[u] + w < dist[v]:
//	    dist[v] = dist[u] + w
//
// After V-1 passes, one more improvement means there is a negative cycle.
package advancedgraphs

const Infinity = int(^uint(0) >> 1)

type Edge struct {
	From   int
	To     int
	Weight int
}

// BellmanFord computes shortest paths from start and detects negative cycles.
//
// Use when:
//
//	The graph can have negative edge weights.
//
// Dijkstra does not work with negative weights. Bellman-Ford is slower but
// handles them and can detect negative cycles.
//
// Time: O(V*E)
// Space: O(V)
func BellmanFord(vertices int, edges []Edge, start int) ([]int, bool) {
	dist := make([]int, vertices)
	for i := range dist {
		dist[i] = Infinity
	}
	dist[start] = 0

	for i := 0; i < vertices-1; i++ {
		for _, edge := range edges {
			if dist[edge.From] == Infinity {
				continue
			}
			if dist[edge.From]+edge.Weight < dist[edge.To] {
				dist[edge.To] = dist[edge.From] + edge.Weight
			}
		}
	}

	for _, edge := range edges {
		if dist[edge.From] != Infinity && dist[edge.From]+edge.Weight < dist[edge.To] {
			return dist, true
		}
	}

	return dist, false
}
