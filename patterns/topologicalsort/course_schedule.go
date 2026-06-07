// Package topologicalsort contains dependency-ordering patterns.
package topologicalsort

// CanFinish returns true if all courses can be completed.
//
// Pattern:
//
//	Build graph and indegree counts.
//	Repeatedly take nodes with indegree 0.
//	If all nodes are processed, there is no cycle.
//
// Time: O(V + E)
// Space: O(V + E)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, edge := range prerequisites {
		course, prereq := edge[0], edge[1]
		graph[prereq] = append(graph[prereq], course)
		indegree[course]++
	}

	var queue []int
	for course, degree := range indegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	seen := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		seen++

		for _, next := range graph[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return seen == numCourses
}
