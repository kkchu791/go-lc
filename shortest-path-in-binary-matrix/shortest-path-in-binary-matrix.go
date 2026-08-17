package shortestpathinbinarymatrix

type Cell struct {
	x int
	y int
	d int
}

type Direction struct {
	rowD int
	colD int
}

func shortestPathBinaryMatrix(grid [][]int) int {

	if grid[0][0] == 1 {
		return -1
	}

	initCell := Cell{x: 0, y: 0, d: 1}

	// create a list with all max possible numbers from the cell
	queue := make([]Cell, 0, len(grid)*len(grid[0]))
	queue = append(queue, initCell)
	head := 0
	seen := make(map[int]struct{})
	targetX := len(grid) - 1
	targetY := len(grid[0]) - 1

	if initCell.x == targetX && initCell.y == targetY {
		return initCell.d
	}

	ad := []Direction{
		{rowD: -1, colD: 0},  // top
		{rowD: -1, colD: 1},  // top right
		{rowD: 0, colD: 1},   // right
		{rowD: 1, colD: 1},   // bottom right
		{rowD: 1, colD: 0},   // bottom
		{rowD: 1, colD: -1},  // bottom left
		{rowD: 0, colD: -1},  // left
		{rowD: -1, colD: -1}, // top left
	}

	for head < len(queue) {
		cc := queue[head] // get first el // [0, 0]
		head++            // move up one
		for _, d := range ad {
			nc := Cell{
				x: cc.x + d.rowD,
				y: cc.y + d.colD,
				d: cc.d + 1,
			}

			isOut := nc.x < 0 || nc.x >= len(grid) || nc.y < 0 || nc.y >= len(grid[0])

			intNC := nc.x*len(grid[0]) + nc.y
			// 0 * 3 + 1 = 1
			// 0 * 3 + 2 = 2
			// 1 * 3 + 1 = 4
			// 1 * 3 + 2 = 5

			_, inSeen := seen[intNC]

			if isOut {
				continue
			}

			if inSeen {
				continue
			}

			if grid[nc.x][nc.y] == 1 {
				continue
			}

			if nc.x == targetX && nc.y == targetY {
				return nc.d
			}

			queue = append(queue, nc)
			seen[intNC] = struct{}{}
		}
	}

	return -1
}
