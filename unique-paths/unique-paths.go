package uniquepaths

func uniquePaths(m int, n int) int {
	type Cell struct {
		r int
		c int
	}
	tX := m - 1
	tY := n - 1
	memo := make(map[Cell]int)
	var explore func(int, int) int

	explore = func(r int, c int) int {

		numOfPaths := 0

		isOut := r < 0 || r >= m || c < 0 || c >= n
		if isOut {
			return 0
		}

		if r == tX && c == tY {
			return 1
		}

		if cNumOfPaths, exists := memo[Cell{r: r, c: c}]; exists {
			return cNumOfPaths
		}

		numOfPaths += explore(r, c+1) // right
		numOfPaths += explore(r+1, c) // bottom

		memo[Cell{r: r, c: c}] = numOfPaths

		return numOfPaths
	}

	return explore(0, 0)
}
