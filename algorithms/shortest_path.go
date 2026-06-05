// CONCEPT 7: Shortest Path
//
// Dijkstra's algorithm finds shortest paths in a weighted graph with
// non-negative edge weights.
//
// This example uses a simple O(V^2 + E) implementation for readability.
// Production code often uses a priority queue for O((V+E) log V).

package algorithms

import "fmt"

type WeightedEdge struct {
	To     string
	Weight int
}

type WeightedGraph map[string][]WeightedEdge

func Dijkstra(graph WeightedGraph, start string) map[string]int {
	const infinity = int(^uint(0) >> 1)

	dist := make(map[string]int)
	visited := make(map[string]bool)

	for node := range graph {
		dist[node] = infinity
	}
	dist[start] = 0

	for range graph {
		node := closestUnvisited(dist, visited)
		if node == "" {
			break
		}
		visited[node] = true

		for _, edge := range graph[node] {
			if dist[node]+edge.Weight < dist[edge.To] {
				dist[edge.To] = dist[node] + edge.Weight
			}
		}
	}

	return dist
}

func closestUnvisited(dist map[string]int, visited map[string]bool) string {
	const infinity = int(^uint(0) >> 1)
	bestNode := ""
	bestDistance := infinity

	for node, distance := range dist {
		if !visited[node] && distance < bestDistance {
			bestNode = node
			bestDistance = distance
		}
	}

	return bestNode
}

func RunShortestPathDemo() {
	graph := WeightedGraph{
		"A": {{To: "B", Weight: 4}, {To: "C", Weight: 2}},
		"B": {{To: "C", Weight: 1}, {To: "D", Weight: 5}},
		"C": {{To: "D", Weight: 8}, {To: "E", Weight: 10}},
		"D": {{To: "E", Weight: 2}},
		"E": {},
	}

	fmt.Printf("shortest distances from A => %v\n", Dijkstra(graph, "A"))
}
