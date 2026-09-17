package numberofprovinces

func findCircleNum(isConnected [][]int) int {
	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	al := make(map[int][]int, 0)

	for r := 0; r < len(isConnected); r++ {
		for c := 0; c < len(isConnected[0]); c++ {
			cc := isConnected[r][c]

			if cc == 1 {
				al[r] = append(al[r], c)
			}
		}
	}

	var dfs func(int)

	dfs = func(node int) {
		if visited[node] {
			return
		}

		if visiting[node] {
			return
		}

		visiting[node] = true

		for _, neigh := range al[node] {
			dfs(neigh)
		}

		delete(visiting, node)
		visited[node] = true
		return
	}

	count := 0
	for node := range isConnected {
		if !visited[node] {
			dfs(node)
			count++
		}
	}

	return count
}
