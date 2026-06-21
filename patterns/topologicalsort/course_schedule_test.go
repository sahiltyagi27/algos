package topologicalsort

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "no cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "diamond dependencies",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Fatalf("CanFinish() = %v, want %v", got, tt.want)
			}

			gotDFS := CanFinishDFS(tt.numCourses, tt.prerequisites)
			if gotDFS != tt.want {
				t.Fatalf("CanFinishDFS() = %v, want %v", gotDFS, tt.want)
			}
		})
	}
}

func TestFindOrder(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	order := FindOrder(numCourses, prerequisites)
	if len(order) != numCourses {
		t.Fatalf("FindOrder() length = %d, want %d", len(order), numCourses)
	}

	position := make(map[int]int)
	for i, course := range order {
		position[course] = i
	}

	for _, edge := range prerequisites {
		course, prereq := edge[0], edge[1]
		if position[prereq] > position[course] {
			t.Fatalf("invalid order %v: prerequisite %d appears after course %d", order, prereq, course)
		}
	}
}

func TestFindOrderCycle(t *testing.T) {
	order := FindOrder(2, [][]int{{1, 0}, {0, 1}})
	if order != nil {
		t.Fatalf("FindOrder() = %v, want nil for cycle", order)
	}
}
