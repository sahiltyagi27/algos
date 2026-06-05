// CONCEPT 6: Graph Traversal
//
// BFS:
//   Breadth-first search explores level by level using a queue.
//
// DFS:
//   Depth-first search explores deeply using recursion or a stack.

package algorithms

import "fmt"

type Graph map[string][]string

func BFS(graph Graph, start string) []string {
	visited := map[string]bool{start: true}
	queue := []string{start}
	var order []string

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for _, next := range graph[node] {
			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, next)
		}
	}

	return order
}

func DFS(graph Graph, start string) []string {
	visited := make(map[string]bool)
	var order []string

	var visit func(string)
	visit = func(node string) {
		if visited[node] {
			return
		}
		visited[node] = true
		order = append(order, node)
		for _, next := range graph[node] {
			visit(next)
		}
	}

	visit(start)
	return order
}

func RunGraphTraversalDemo() {
	graph := Graph{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"E"},
		"D": {},
		"E": {},
	}

	fmt.Printf("BFS from A => %v\n", BFS(graph, "A"))
	fmt.Printf("DFS from A => %v\n", DFS(graph, "A"))
}
