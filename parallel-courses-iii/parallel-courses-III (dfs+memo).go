package parallelcoursesiii

func minimumTimeDFS(n int, relations [][]int, time []int) int {
	al := make(map[int][]int, 0)
	res := 0
	memo := make([]int, n)

	for _, edge := range relations {
		al[edge[0]-1] = append(al[edge[0]-1], edge[1]-1)
	}

	var dfs func(int) int
	dfs = func(node int) int {
		if memo[node] > 0 {
			return memo[node]
		}

		ans := 0

		for _, neigh := range al[node] {
			ans = max(ans, dfs(neigh))
		}

		maxAtThisPoint := time[node] + ans
		memo[node] += maxAtThisPoint
		return memo[node]
	}

	for node := range n {
		res = max(res, dfs(node))
	}

	return res
}
