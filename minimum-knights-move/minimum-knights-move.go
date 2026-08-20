package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minKnightsMove(x, y int) int {
	x = Abs(x)
	y = Abs(y)

	type Cell struct {
		r int
		c int
		d int
	}

	src := Cell{r: 0, c: 0, d: 0}
	queue := []Cell{src}
	seen := make(map[string]struct{})

	type Direction struct {
		r int
		c int
	}

	allDirections := []Direction{
		{r: -2, c: 1},  // top right
		{r: -1, c: 2},  // right up
		{r: 1, c: 2},   // right down
		{r: 2, c: 1},   // bottom right
		{r: 2, c: -1},  // bottom left
		{r: 1, c: -2},  //  left down
		{r: -1, c: -2}, // left up
		{r: -2, c: -1}, // top left
	}

	for len(queue) > 0 {
		// Shift operation
		cc := queue[0]
		queue = queue[1:]

		for _, direction := range allDirections {

			nc := Cell{
				r: cc.r + direction.r,
				c: cc.c + direction.c,
				d: cc.d + 1,
			}

			// hits out target

			if nc.r == x && nc.c == y {
				return nc.d
			}

			// possibly no upper bounds // there is no out of bounds
			// isOut := nc.r < 0 || nc.c < 0
			// if isOut {
			// 	continue
			// }

			// add discovered cell into our queue
			// but it can't be in seen
			ncString := fmt.Sprint("%,%", nc.r, nc.c)
			if _, exists := seen[ncString]; !exists {
				queue = append(queue, nc)
			}

			seen[ncString] = struct{}{}
		}
	}

	return -1
}

func main() {
	fmt.Println(minKnightsMove(5, 5))
}
