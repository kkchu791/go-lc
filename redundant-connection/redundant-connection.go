package redundantconnection

func findRedundantConnection(edges [][]int) []int {
	al := make(map[int][]int, 0)

	var dfs func(int, []bool)

	dfs = func(node int, visited []bool) {
		if visited[node] {
			return
		}

		visited[node] = true

		for _, neigh := range al[node] {
			dfs(neigh, visited)
		}
	}

	for _, edge := range edges {
		u := edge[0] - 1
		v := edge[1] - 1

		visited := make([]bool, len(edges))

		dfs(u, visited)

		if visited[v] {
			return edge
		}

		al[u] = append(al[u], v)
		al[v] = append(al[v], u)
	}

	return nil
}
