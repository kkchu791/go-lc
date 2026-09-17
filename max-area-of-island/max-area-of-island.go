package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	var dfs func(int, int) int
	dfs = func(r, c int) int {
		isOut := r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
		if isOut {
			return 0
		}

		if grid[r][c] == 0 {
			return 0
		}

		grid[r][c] = 0

		return (dfs(r-1, c) + //0
			dfs(r, c+1) + //0
			dfs(r+1, c) + // 0
			dfs(r, c-1) + 1) // 0 + 1
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				maxArea = max(maxArea, dfs(r, c))
			}
		}
	}

	return maxArea
}
