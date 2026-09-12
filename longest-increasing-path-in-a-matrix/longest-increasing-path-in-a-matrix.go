package longestincreasingpath

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	memo := make(map[string]int, 0)
	matrixMax := 0

	var gmp func(int, int, int) int

	gmp = func(r, c, prevVal int) int {

		isOut := r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[0])
		if isOut {
			return 0
		}

		if matrix[r][c] <= prevVal {
			return 0
		}

		memoKey := fmt.Sprintf("%d,%d", r, c)
		if max, exists := memo[memoKey]; exists {
			return max
		}

		max := max(
			gmp(r-1, c, matrix[r][c]), // top
			gmp(r, c+1, matrix[r][c]), // right
			gmp(r+1, c, matrix[r][c]), // bottom
			gmp(r, c-1, matrix[r][c]), // left
		) + 1

		memo[memoKey] = max
		return max
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			matrixMax = max(matrixMax, gmp(r, c, -1))
		}
	}

	return matrixMax
}
