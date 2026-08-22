package coursescheduleii

import "slices"

func findOrder2(numCourses int, prerequisites [][]int) []int {
	visited := make(map[int]struct{})
	visiting := make(map[int]struct{})
	res := make([]int, 0)

	al := make(map[int][]int)
	for _, prereq := range prerequisites {
		al[prereq[1]] = append(al[prereq[1]], prereq[0])
	}

	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if _, exists := visited[course]; exists {
			return false
		}

		if _, exists := visiting[course]; exists {
			return true
		}

		visiting[course] = struct{}{}

		for _, nc := range al[course] {
			if hasCycle(nc) {
				return true
			}
		}

		res = append(res, course)

		visited[course] = struct{}{}
		delete(visiting, course)

		return false
	}

	for course := range numCourses {
		if hasCycle(course) {
			return []int{}
		}
	}

	slices.Reverse(res)
	return res
}
