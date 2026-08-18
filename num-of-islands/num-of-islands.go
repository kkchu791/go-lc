package numofislands

func numIslands(grid [][]byte) int {
	count := 0

	type Position struct {
		r int
		c int
	}

	memo := make(map[Position]struct{})

	var dfs func(r, c int)

	dfs = func(r, c int) {
		isOut := r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
		if isOut {
			return
		}

		if grid[r][c] == '0' {
			return
		}

		if _, exists := memo[Position{r, c}]; exists {
			return
		}

		memo[Position{r, c}] = struct{}{}

		dfs(r-1, c) // top
		dfs(r, c+1) // right
		dfs(r+1, c) // bottom
		dfs(r, c-1) // left
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if _, exists := memo[Position{r, c}]; grid[r][c] == '1' && !exists {
				dfs(r, c)
				count += 1
			}
		}
	}

	return count
}
