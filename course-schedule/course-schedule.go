package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	// assumption: if course has no preReqs,
	// no cycle possible if you hit that number,
	// so no need to add to our adjList

	adjList := make(map[int][]int)
	for _, req := range prerequisites {
		prs, exists := adjList[req[1]]

		if exists {
			prs = append(prs, req[0])
			adjList[req[1]] = prs
		} else {
			list := make([]int, 0)
			adjList[req[1]] = append(list, req[0])
		}
	}

	var hasCycle func(int) bool
	visited := make(map[int]struct{})
	visiting := make(map[int]struct{})

	hasCycle = func(course int) bool {
		if _, exists := visited[course]; exists {
			return false
		}

		if _, exists := visiting[course]; exists {
			return true
		}

		visiting[course] = struct{}{}

		if pReqs, exists := adjList[course]; exists {
			for _, preReq := range pReqs {
				if hasCycle(preReq) {
					return true
				}
			}
		}

		visited[course] = struct{}{}
		delete(visiting, course)

		return false
	}

	for course := range numCourses {
		if hasCycle(course) {
			return false
		}
	}
	return true
}
