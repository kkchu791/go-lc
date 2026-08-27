package validtree

func validTree(n int, edges [][]int) bool {
	al := make(map[int][]int, n)
	visiting := make(map[int]struct{})
	visited := make(map[int]struct{})

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
	}

	var hasCycle func(int, int) bool

	hasCycle = func(node, parent int) bool {
		if _, exists := visiting[node]; exists {
			return true
		}

		if _, exists := visited[node]; exists {
			return false
		}

		visiting[node] = struct{}{}

		for _, neighbor := range al[node] {
			if parent != neighbor {
				if hasCycle(neighbor, node) {
					return true
				}
			}
		}

		visited[node] = struct{}{}
		delete(visiting, node)
		return false
	}

	if hasCycle(0, -1) {
		return false
	}

	return n == len(visited)
}
