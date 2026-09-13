package rottingoranges

import "fmt"

func orangesRotting(grid [][]int) int {
	rottenSet := make(map[string]struct{})
	freshCount := 0
	rottenCount := 0
	queue := make([][2]int, 0)
	round := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			cc := grid[r][c]

			if cc == 1 {
				freshCount++
			}

			key := fmt.Sprintf("%d,%d", r, c)
			if cc == 2 {
				rottenSet[key] = struct{}{}
				coordinates := [2]int{r, c}
				queue = append(queue, coordinates)
				rottenCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	helperCount := len(queue)

	for len(queue) > 0 {
		fmt.Println(queue)
		cc := queue[0]
		queue = queue[1:]

		row := cc[0]
		col := cc[1]

		for _, d := range directions {
			r := row + d[0]
			c := col + d[1]

			isOut := r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
			if isOut {
				continue
			}

			key := fmt.Sprintf("%d,%d", r, c)
			if _, exists := rottenSet[key]; exists {
				continue
			}

			if grid[r][c] == 0 {
				continue
			}

			// it has to be 1
			rottenSet[key] = struct{}{}
			coordinate := [2]int{r, c}
			queue = append(queue, coordinate)
		}

		helperCount--

		if helperCount == 0 {
			helperCount = len(queue)
			round++
		}
	}

	if len(rottenSet) == freshCount+rottenCount {
		return round - 1 // to make room for minute 0
	} else {
		return -1
	}
}
