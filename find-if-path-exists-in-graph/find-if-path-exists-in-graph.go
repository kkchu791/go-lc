package findifpathexistsingraph

func validPath(n int, edges [][]int, source int, destination int) bool {
	// build al
	al := make(map[int][]int, 0)
	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
	}

	var dfs func(int, int) bool
	visited := make(map[int]bool, 0)
	dfs = func(node, parent int) bool {
		if visited[node] {
			return false
		}

		if node == destination {
			return true
		}

		visited[node] = true

		for _, neigh := range al[node] {
			if parent == neigh {
				continue
			}

			if dfs(neigh, node) {
				return true
			}
		}

		return false
	}

	return dfs(source, -1)
}
