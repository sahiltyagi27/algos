// Package topologicalsort contains dependency-ordering patterns.
package topologicalsort

// CanFinish returns true if all courses can be completed.
//
// Course Schedule input uses pairs like [course, prerequisite].
// Example: [1, 0] means "take course 0 before course 1".
// So the graph edge is prerequisite -> course.
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

// CanFinishDFS returns true if all courses can be completed using DFS cycle detection.
//
// State meaning:
//
//	0 = unvisited
//	1 = visiting, currently in recursion stack
//	2 = visited, already proven safe
//
// If DFS reaches a node with state 1, there is a cycle.
//
// Time: O(V + E)
// Space: O(V + E)
func CanFinishDFS(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, edge := range prerequisites {
		course, prereq := edge[0], edge[1]
		graph[prereq] = append(graph[prereq], course)
	}

	state := make([]int, numCourses)
	var hasCycle func(course int) bool

	hasCycle = func(course int) bool {
		if state[course] == 1 {
			return true
		}
		if state[course] == 2 {
			return false
		}

		state[course] = 1
		for _, next := range graph[course] {
			if hasCycle(next) {
				return true
			}
		}
		state[course] = 2
		return false
	}

	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}

	return true
}

// FindOrder returns one valid course order.
//
// This is the Course Schedule II version of the same pattern.
// If there is a cycle, no valid ordering exists, so it returns nil.
//
// Time: O(V + E)
// Space: O(V + E)
func FindOrder(numCourses int, prerequisites [][]int) []int {
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

	order := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)

		for _, next := range graph[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(order) != numCourses {
		return nil
	}
	return order
}
