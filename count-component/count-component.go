package countcomponent

func countComponents(n int, edges [][]int) int {
	visited := make([]bool, n)
	count := 0
	al := make(map[int][]int)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
	}

	var dfs func(int, int)

	dfs = func(v, pi int) {
		visited[v] = true

		for _, neigh := range al[v] {
			if pi == neigh {
				continue
			}

			if visited[neigh] {
				continue
			}

			dfs(neigh, v)
		}
	}

	for v := range n {
		if !visited[v] {
			dfs(v, -1)
			count++
		}
	}

	return count
}
