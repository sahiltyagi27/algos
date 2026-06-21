# Topological Sort: Course Schedule

Course Schedule is a dependency-ordering graph problem.

Common prompt:

> There are `numCourses` courses and prerequisites like `[course, prerequisite]`. Return whether all courses can be completed.

Example:

```text
numCourses = 2
prerequisites = [[1, 0]]
```

Meaning:

```text
Take course 0 before course 1.
```

So the graph edge is:

```text
0 -> 1
```

## Approach 1: BFS / Kahn's Algorithm

If a course has no prerequisites, it can be taken immediately.

That is exactly what `indegree == 0` means.

```text
indegree[course] = number of prerequisites still remaining
```

Algorithm:

```text
1. Build adjacency list.
2. Build indegree array.
3. Put all courses with indegree 0 into a queue.
4. Pop courses from queue one by one.
5. For each popped course, reduce indegree of dependent courses.
6. If a dependent course reaches indegree 0, push it into queue.
7. If processed count == numCourses, no cycle exists.
8. Otherwise, a cycle exists and some courses cannot be completed.
```

## Why This Detects A Cycle

In a cycle, every course inside the cycle depends on another course inside the same cycle.

Example:

```text
0 -> 1
1 -> 0
```

Both courses have indegree `1`.

There is no course with indegree `0`, so the queue becomes empty before all courses are processed.

That means:

```text
processed count < numCourses
```

So the answer is false.

## Approach 2: DFS Cycle Detection

Course Schedule can also be solved with DFS.

DFS state:

```text
0 = unvisited
1 = visiting, currently in recursion stack
2 = visited, already proven safe
```

Cycle rule:

```text
If DFS reaches a node that is already "visiting", there is a cycle.
```

Why?

```text
visiting means this course is already in the current dependency path.
If we reach it again before finishing that path, dependencies loop back.
```

Example cycle:

```text
0 -> 1
1 -> 0
```

DFS from `0`:

```text
mark 0 visiting
go to 1
mark 1 visiting
go to 0
0 is already visiting -> cycle
```

DFS algorithm:

```text
1. Build adjacency list.
2. Keep state array for all courses.
3. DFS every unvisited course.
4. Mark course as visiting before exploring neighbors.
5. If a neighbor is visiting, cycle exists.
6. Mark course as visited after all neighbors are safe.
7. If no cycle is found, all courses can be completed.
```

## Dry Run

Input:

```text
numCourses = 4
prerequisites = [[1,0], [2,0], [3,1], [3,2]]
```

Meaning:

```text
0 -> 1
0 -> 2
1 -> 3
2 -> 3
```

Indegree:

```text
course 0 = 0
course 1 = 1
course 2 = 1
course 3 = 2
```

Queue starts with:

```text
[0]
```

Process:

```text
pop 0 -> reduce indegree of 1 and 2
course 1 becomes 0 -> push 1
course 2 becomes 0 -> push 2

pop 1 -> reduce indegree of 3 from 2 to 1
pop 2 -> reduce indegree of 3 from 1 to 0 -> push 3
pop 3
```

Processed all 4 courses:

```text
true
```

One valid order:

```text
[0, 1, 2, 3]
```

`[0, 2, 1, 3]` is also valid.

## Code

See:

- [course_schedule.go](course_schedule.go)

Functions:

```text
CanFinish    -> Course Schedule I using BFS / Kahn's Algorithm
CanFinishDFS -> Course Schedule I using DFS cycle detection
FindOrder    -> Course Schedule II, returns one valid order or nil
```

## Interview Explanation

> I model courses as a directed graph. For each prerequisite pair `[course, prerequisite]`, I add an edge from prerequisite to course. Then I calculate indegree for every course, where indegree means how many prerequisites are still needed. I start with all courses that have indegree 0 because they can be taken immediately. Every time I process a course, I reduce the indegree of courses depending on it. If a course's indegree becomes 0, I add it to the queue. If I process all courses, there is no cycle and the schedule is possible. If not, a cycle exists.

DFS interview explanation:

> Another way is DFS cycle detection. I keep three states: unvisited, visiting and visited. When DFS enters a course, I mark it as visiting. If I reach another course that is already visiting, that means the dependency path loops back, so there is a cycle. After all neighbors are safe, I mark the course as visited. If DFS finishes without finding a cycle, all courses can be completed.

## Common Mistakes

```text
Reversing the edge direction.
Forgetting that [1,0] means 0 -> 1.
Using DFS but not tracking visiting/visited states.
Returning true just because the queue is not empty initially.
Forgetting disconnected courses with no prerequisites.
```

## Complexity

```text
V = number of courses
E = number of prerequisite pairs

Time:  O(V + E)
Space: O(V + E)
```

## When To Use This Pattern

Use topological sort when the problem has:

```text
dependencies
ordering constraints
prerequisites
build order
task scheduling
cycle detection in directed graph
```
