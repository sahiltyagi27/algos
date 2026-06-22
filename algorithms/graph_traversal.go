// CONCEPT 6: Graph Traversal
//
// BFS:
//   Breadth-first search explores level by level using a queue.
//
// Visual:
//
//	      A
//	    /   \
//	   B     C
//	  /       \
//	 D         E
//
// BFS order:
//
//	A, B, C, D, E
//
// DFS:
//   Depth-first search explores deeply using recursion or a stack.
//
// DFS order can be:
//
//	A, B, D, C, E

package algorithms

import "fmt"

type Graph map[string][]string

// BFS uses a queue.
//
// Queue flow:
//
//	enqueue start
//	pop front
//	visit node
//	enqueue unseen neighbors
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

// DFS uses recursion here.
//
// Recursion flow:
//
//	visit node
//	recursively visit each unseen neighbor
//	backtrack when no more neighbors exist
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

	grid := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}

	fmt.Printf("surrounded city cells => %d\n", CountSurroundedCities(grid))
	fmt.Printf("surrounded city components => %d\n", CountSurroundedCityComponents(grid))
}

// CountSurroundedCities counts individual 1-cells that have 0 on all four
// sides.
//
// Check directions:
//
//	    up
//	left X right
//	   down
//
// This is not component DFS. It checks each cell independently.
func CountSurroundedCities(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 {
				continue
			}

			surrounded := true
			for _, d := range dirs {
				nr := r + d[0]
				nc := c + d[1]

				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != 0 {
					surrounded = false
					break
				}
			}

			if surrounded {
				count++
			}
		}
	}

	return count
}

// CountSurroundedCityComponents counts connected groups of 1s that do not touch
// the boundary.
//
// Here DFS explores an entire component:
//
//	1 1 0
//	0 1 0
//
// If any cell in the component touches boundary, the whole component is not
// surrounded.
func CountSurroundedCityComponents(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		visited[r][c] = true
		isSurrounded := true

		if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
			isSurrounded = false
		}

		for _, d := range dirs {
			nr := r + d[0]
			nc := c + d[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				isSurrounded = false
				continue
			}

			if grid[nr][nc] == 0 {
				continue
			}

			if grid[nr][nc] == 1 && !visited[nr][nc] {
				if !dfs(nr, nc) {
					isSurrounded = false
				}
			}
		}

		return isSurrounded
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 && !visited[r][c] && dfs(r, c) {
				count++
			}
		}
	}

	return count
}
