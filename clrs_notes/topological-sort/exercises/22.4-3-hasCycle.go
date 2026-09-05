package exercises

func detectCycle(n int, edges [][]int) bool {
	visited := make([]bool, n)
	al := make(map[int][]int, 0)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
	}

	var hasCycle func(int, int) bool
	hasCycle = func(u, parent int) bool {
		visited[u] = true

		for _, neigh := range al[u] {
			if parent == neigh {
				continue
			}

			if visited[neigh] {
				return true
			}

			if hasCycle(neigh, u) {
				return true
			}
		}

		return false
	}

	for node := range n {
		if !visited[node] && hasCycle(node, -1) { // remember not to revisit again in the outer loop if alreaded visited
			return true
		}
	}

	return false
}
