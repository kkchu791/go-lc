package exercises

func countSimplePath() int {
	edges := [][]string{
		{"A", "D"},
		{"A", "B"},
		{"A", "C"},
		{"D", "B"},
		{"B", "C"},
	}

	source := "A"
	target := "C"

	al := make(map[string][]string, 0)
	memo := make(map[string]int)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
	}

	var dfs func(string) int

	dfs = func(node string) int {
		pathCount := 0

		if _, exists := memo[node]; exists {
			return memo[node]
		}

		if node == target {
			return 1
		}

		for _, neigh := range al[node] {
			pathCount += dfs(neigh)
		}

		memo[node] = pathCount
		return memo[node]

	}

	return dfs(source)
}
