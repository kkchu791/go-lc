package pacificatlanticwaterflow

func pacificAtlantic(heights [][]int) [][]int {
	// lets create the matrix memos firs
	pMemo := make([][]bool, len(heights))
	aMemo := make([][]bool, len(heights))
	res := make([][]int, 0)

	for r := 0; r < len(heights); r++ {
		pMemo[r] = make([]bool, len(heights[0]))
		aMemo[r] = make([]bool, len(heights[0]))
	}

	var dfs func(int, int, int, [][]bool)

	dfs = func(r, c, prevHeight int, grid [][]bool) {

		isOut := r < 0 || r >= len(heights) || c < 0 || c >= len(heights[0])
		if isOut {
			return
		}

		if grid[r][c] == true {
			return
		}

		if heights[r][c] < prevHeight {
			return
		}

		grid[r][c] = true

		dfs(r-1, c, heights[r][c], grid) //top
		dfs(r, c+1, heights[r][c], grid) //right
		dfs(r+1, c, heights[r][c], grid) //bottom
		dfs(r, c-1, heights[r][c], grid) //left
	}

	// dfs the top row for pacific
	// dfs the bottom row for atlantic
	for c := 0; c < len(heights[0]); c++ {
		dfs(0, c, -1, pMemo)
		dfs(len(heights)-1, c, -1, aMemo)
	}

	// dfs the left col for pacific
	// dfs the right col for atlantic
	for r := 0; r < len(heights); r++ {
		dfs(r, 0, -1, pMemo)
		dfs(r, len(heights[0])-1, -1, aMemo)
	}

	for r := 0; r < len(heights); r++ {
		for c := 0; c < len(heights[0]); c++ {

			if pMemo[r][c] == true && aMemo[r][c] == true {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
